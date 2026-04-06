package ws

import (
	"encoding/json"
	"log"

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

	responseChat := c.processChat(reqChat.Content, reqChat.RoomID, *c.userID)

	broadCast := &BroadCast{
		Type:    values.MessageTypeChat,
		RoomID:  reqChat.RoomID,
		Content: responseChat,
	}
	c.Hub.Broadcast <- broadCast

	go c.ollamaHandlerAsync(reqChat)
}

func (c *Client) ollamaHandlerAsync(reqChat domain.ChatMessage) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ollama panic:", r)
		}
	}()

	c.ollamaHandler(reqChat)
}

func (c *Client) ollamaHandler(reqChat domain.ChatMessage) {
	reqData := ollamaDomain.RequestChat{
		Role:    core_values.OllamaRoleUser,
		Content: reqChat.Content,
	}
	roomHistory, err := c.svc.FindHistoryByRoomID(reqChat.RoomID, core_values.ChatHistoryLimit)
	if err != nil {
		log.Println(err.Error())
		return
	}
	aiResponse := c.ollamaSvc.Chat(reqData, *roomHistory)

	responseChat := c.processChat(aiResponse.Message.Content, reqChat.RoomID, default_data.GetAIWorker().ID)
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
			Body:   aiResponse.Message.Content,
		}
		c.pushSvc.Push(reqPush)
	}
}
