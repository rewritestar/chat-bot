package ws

import (
	"encoding/json"
	"errors"
	"log"

	chatDomain "chat-bot/src/common-service/chat/domain"
	chatMessage "chat-bot/src/common-service/chat/message"
	core_values "chat-bot/src/core/values"
	"chat-bot/src/core/ws/domain"
	"chat-bot/src/core/ws/values"
	"chat-bot/src/initial/default_data"
)

func (c *Client) readHandler(message []byte) {
	reqMsg := domain.Message{}
	if err := json.Unmarshal(message, &reqMsg); err != nil {
		log.Printf("error: %v", err)
		return
	}
	dataBytes, err := json.Marshal(reqMsg.Data)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	switch reqMsg.Type {
	case values.MessageTypeAuth:
		c.authHandler(dataBytes)
	case values.MessageTypeChat:
		c.chatHandler(dataBytes)
	default:
	}
}

func (c *Client) authHandler(data []byte) {
	reqAuth := domain.AuthMessage{}
	json.Unmarshal(data, &reqAuth)

	if err := c.authorizeWsJwt(reqAuth.Token); err != nil {
		log.Printf("error: %v", err)
		return
	}
}

func (c *Client) chatHandler(data []byte) {
	reqChat := domain.ChatMessage{}
	json.Unmarshal(data, &reqChat)

	userID, ok := c.ctx.Get(core_values.WorkerIDKey)
	if !ok {
		err := errors.New("worker id does not exist.")
		log.Printf("error: %v", err)
		return
	}
	roomHistory, err := c.svc.FindHistoryByRoomID(reqChat.RoomID)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	domainChat := chatDomain.Chat{}
	domainChat.Content = reqChat.Content
	domainChat.RoomID = reqChat.RoomID
	domainChat.CreatorID = userID.(uint)

	savedChat, err := c.svc.SaveChat(domainChat)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	responseChat := chatMessage.ResponseChat{}
	responseChat.Build(savedChat)

	chatBytes, err := json.Marshal(responseChat)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	c.Hub.broadcast <- chatBytes
	c.ollamaHandler(reqChat, *roomHistory)
}

func (c *Client) ollamaHandler(reqChat domain.ChatMessage, roomHistory chatDomain.Room) {
	response := c.ollamaSvc.Chat(reqChat.Content, roomHistory)

	domainChat := chatDomain.Chat{}
	domainChat.Content = response.Message.Content
	domainChat.RoomID = reqChat.RoomID
	domainChat.CreatorID = default_data.GetAIWorker().ID

	savedChat, err := c.svc.SaveChat(domainChat)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	responseChat := chatMessage.ResponseChat{}
	responseChat.Build(savedChat)

	chatBytes, err := json.Marshal(responseChat)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	c.Hub.broadcast <- chatBytes
}
