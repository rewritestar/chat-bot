package service

import (
	"chat-bot/src/app/ws/values"
	"chat-bot/src/common-service/chat"
	"chat-bot/src/core/ws"

	"github.com/gin-gonic/gin"
)

type wsService struct {
	chatService chat.ChatService
}

func NewWsService(chatService chat.ChatService) WsService {
	return &wsService{
		chatService,
	}
}

func (s *wsService) AddClient(ctx *gin.Context) error {
	conn, err := values.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return err
	}
	client := ws.NewClient(ws.GetHub(), conn, make(chan []byte, 256), ctx, s.chatService)
	client.Hub.Register <- client

	go client.WritePump()
	go client.ReadPump()
	return nil
}
