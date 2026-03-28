package aischeduler

import (
	"log"
	"math/rand"
	"time"

	"chat-bot/src/common-service/chat"
	chatDomain "chat-bot/src/common-service/chat/domain"
	chatMessage "chat-bot/src/common-service/chat/message"
	"chat-bot/src/core/ollama"
	ollamaDomain "chat-bot/src/core/ollama/domain"
	core_values "chat-bot/src/core/values"
	"chat-bot/src/core/ws"
	"chat-bot/src/initial/ai_scheduler/values"
	"chat-bot/src/initial/default_data"
)

type aiScheduler struct {
	ollamaSvc ollama.OllamaService
	chatSvc   chat.ChatService
	hub       *ws.Hub
}

type AiScheduler interface {
	StartAIScheduler()
}

func NewAiSchduler(ollamaSvc ollama.OllamaService, chatSvc chat.ChatService, hub *ws.Hub) AiScheduler {
	return &aiScheduler{
		ollamaSvc,
		chatSvc,
		hub,
	}
}

func (a *aiScheduler) StartAIScheduler() {
	ticker := time.NewTicker(30 * time.Second)

	go func() {
		for range ticker.C {
			a.runAiSchedule()
		}
	}()
}

func (a *aiScheduler) runAiSchedule() {
	roomList, err := a.chatSvc.FindAllRoomTickAiSchedule()
	if err != nil {
		return
	}

	for _, room := range *roomList {
		if len(room.ChatList) == 0 {
			a.generateAiTalk(room.ID)
		} else if len(room.ChatList) > 0 {
			lastChat := room.ChatList[0]
			if lastChat.DateCreated.Add(time.Hour).Before(time.Now()) {
				a.generateAiTalk(room.ID)
			}
		}

		next := nextRandomTime(2, 6)
		room.AiProactiveNextDate = &next
		a.chatSvc.UpdateRoom(room)
	}
}

func (a *aiScheduler) generateAiTalk(roomID uint) {
	reqData := ollamaDomain.RequestChat{
		Role:    core_values.OllamaRoleSystem,
		Content: values.ProactiveSystem,
	}
	roomHistory, err := a.chatSvc.FindHistoryByRoomID(roomID)
	if err != nil {
		log.Println(err.Error())
		return
	}

	response := a.ollamaSvc.Chat(reqData, *roomHistory)

	domainChat := chatDomain.Chat{}
	domainChat.Content = response.Message.Content
	domainChat.RoomID = roomID
	domainChat.CreatorID = default_data.GetAIWorker().ID
	savedChat, err := a.chatSvc.SaveChat(domainChat)
	if err != nil {
		log.Println(err.Error())
		return
	}

	responseChat := chatMessage.ResponseChat{}
	responseChat.Build(savedChat)
	broadCast := &ws.BroadCast{
		RoomID:  roomID,
		Content: responseChat,
	}
	a.hub.Broadcast <- broadCast
}

func nextRandomTime(minHours, maxHours int) time.Time {
	now := time.Now()
	diff := rand.Intn((maxHours-minHours)*60) + (minHours * 60)
	return now.Add(time.Duration(diff) * time.Minute)
}
