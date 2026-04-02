package ws

import (
	"encoding/json"
	"log"

	chatDomain "chat-bot/src/common-service/chat/domain"
	chatMessage "chat-bot/src/common-service/chat/message"
	ollamaDomain "chat-bot/src/core/ollama/domain"
	core_values "chat-bot/src/core/values"
	webPushDomain "chat-bot/src/core/web_push/domain"
	"chat-bot/src/core/ws/domain"
	"chat-bot/src/core/ws/values"
	"chat-bot/src/initial/default_data"
)

func (c *Client) readHandler(message []byte) {
	reqMsg := domain.Message{}
	err := json.Unmarshal(message, &reqMsg)
	if err != nil {
		log.Println(err.Error())
		return
	}

	var dataBytes []byte
	if reqMsg.Data != nil {
		dataBytes, err = json.Marshal(reqMsg.Data)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}

	switch reqMsg.Type {
	case values.MessageTypeAuth:
		c.authHandler(dataBytes)
	case values.MessageTypeJoin:
		c.joinHandler(dataBytes)
	case values.MessageTypeLeave:
		c.leaveHandler()
	case values.MessageTypeChat:
		c.chatHandler(dataBytes)
	default:
	}
}

func (c *Client) authHandler(data []byte) {
	reqAuth := domain.AuthMessage{}
	json.Unmarshal(data, &reqAuth)

	if err := c.authorizeWsJwt(reqAuth.Token); err != nil {
		log.Println(err.Error())
		c.conn.Close()
		return
	}
	c.Hub.Register <- c

	roomList, err := c.svc.FindRoomByCreatorID(*c.userID)
	if err != nil {
		return
	}
	for _, room := range *roomList {
		joinRoom := &JoinRoom{
			RoomID: room.ID,
			Client: c,
		}
		c.Hub.JoinRoom <- joinRoom
	}
	c.CurrentRoomID = nil
}

func (c *Client) joinHandler(data []byte) {
	reqJoin := domain.JoinMessage{}
	json.Unmarshal(data, &reqJoin)

	joinRoom := &JoinRoom{
		RoomID: reqJoin.RoomID,
		Client: c,
	}
	c.Hub.JoinRoom <- joinRoom
}

func (c *Client) leaveHandler() {
	c.Hub.leaveRoom <- c
}

func (c *Client) chatHandler(data []byte) {
	reqChat := domain.ChatMessage{}
	json.Unmarshal(data, &reqChat)

	roomHistory, err := c.svc.FindHistoryByRoomID(reqChat.RoomID)
	if err != nil {
		log.Println(err.Error())
		return
	}

	domainChat := chatDomain.Chat{}
	domainChat.Content = reqChat.Content
	domainChat.RoomID = reqChat.RoomID
	domainChat.CreatorID = *c.userID

	savedChat, err := c.svc.SaveChat(domainChat)
	if err != nil {
		log.Println(err.Error())
		return
	}

	responseChat := chatMessage.ResponseChat{}
	responseChat.Build(savedChat)

	broadCast := &BroadCast{
		Type:    values.MessageTypeChat,
		RoomID:  reqChat.RoomID,
		Content: responseChat,
	}
	c.Hub.Broadcast <- broadCast
	go c.ollamaHandlerAsync(reqChat, *roomHistory)
}

func (c *Client) ollamaHandlerAsync(reqChat domain.ChatMessage, roomHistory chatDomain.Room) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ollama panic:", r)
		}
	}()

	c.ollamaHandler(reqChat, roomHistory)
}

func (c *Client) ollamaHandler(reqChat domain.ChatMessage, roomHistory chatDomain.Room) {
	reqData := ollamaDomain.RequestChat{
		Role:    core_values.OllamaRoleUser,
		Content: reqChat.Content,
	}
	response := c.ollamaSvc.Chat(reqData, roomHistory)

	domainChat := chatDomain.Chat{}
	domainChat.Content = response.Message.Content
	domainChat.RoomID = reqChat.RoomID
	domainChat.CreatorID = default_data.GetAIWorker().ID

	savedChat, err := c.svc.SaveChat(domainChat)
	if err != nil {
		log.Println(err.Error())
		return
	}

	responseChat := chatMessage.ResponseChat{}
	responseChat.Build(savedChat)

	broadCast := &BroadCast{
		Type:    values.MessageTypeChat,
		RoomID:  reqChat.RoomID,
		Content: responseChat,
	}
	c.Hub.Broadcast <- broadCast

	if c.Hub.IsPushTarget(roomHistory.CreatorID, roomHistory.ID) {
		reqPush := webPushDomain.RequestPush{
			UserID: roomHistory.CreatorID,
			RoomID: reqChat.RoomID,
			Title:  roomHistory.Name,
			Body:   response.Message.Content,
		}
		c.pushSvc.Push(reqPush)
	}
}
