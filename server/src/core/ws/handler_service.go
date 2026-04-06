package ws

import (
	"log"
	"time"

	chatDomain "chat-bot/src/common-service/chat/domain"
	chatMessage "chat-bot/src/common-service/chat/message"
	core_values "chat-bot/src/core/values"
)

func (c *Client) processChat(content string, roomID, creatorID uint) *chatMessage.ResponseChat {
	domainChat := chatDomain.Chat{}
	domainChat.Content = content
	domainChat.RoomID = roomID
	domainChat.CreatorID = creatorID

	savedChat, err := c.saveChat(domainChat)
	if err != nil {
		return nil
	}

	responseChat := chatMessage.ResponseChat{}
	responseChat.Build(savedChat)
	return &responseChat
}

func (c *Client) saveChat(domainChat chatDomain.Chat) (*chatDomain.Chat, error) {
	savedChat, err := c.svc.SaveChat(domainChat)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	updatedRoom, err := c.svc.UpdateMemoryPendingCount(savedChat.RoomID)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	if *updatedRoom.MemoryPendingCount >= uint(core_values.SummaryHistoryLimit) {
		go c.summaryAsync(*updatedRoom)
	}
	return savedChat, nil
}

func (c *Client) summaryAsync(updatedRoom chatDomain.Room) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("summary panic:", r)
		}
	}()

	roomHistory, err := c.svc.FindHistoryByRoomID(updatedRoom.ID, core_values.SummaryHistoryLimit)
	if err != nil {
		log.Println(err.Error())
		return
	}

	response := c.ollamaSvc.Summarize(*roomHistory)
	updatedRoom.MemoryPendingCount = core_values.UintToPointer(uint(0))
	updatedRoom.Memory = response.Message.Content
	updatedRoom.MemoryUpdated = time.Now()

	if _, err := c.svc.UpdateRoom(updatedRoom); err != nil {
		return
	}
}
