package ws

import (
	chatDomain "chat-bot/src/common-service/chat/domain"
	chatMessage "chat-bot/src/common-service/chat/message"
	core_values "chat-bot/src/core/values"
	"chat-bot/src/core/ws/domain"
	"chat-bot/src/core/ws/values"
	"encoding/json"
	"errors"
	"log"
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
	return
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

	domainChat := chatDomain.Chat{}
	domainChat.Content = reqChat.Content
	domainChat.RoomID = reqChat.RoomID
	workerID, ok := c.ctx.Get(core_values.WorkerIDKey)
	if !ok {
		err := errors.New("worker id does not exist.")
		log.Printf("error: %v", err)
		return
	}
	domainChat.CreatorID = workerID.(uint)

	if _, err := c.svc.SaveChat(domainChat); err != nil {
		log.Printf("error: %v", err)
		return
	}

	responseChat := chatMessage.ResponseChat{}
	responseChat.Build(&domainChat)

	chatBytes, err := json.Marshal(responseChat)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	c.Hub.broadcast <- chatBytes

}
