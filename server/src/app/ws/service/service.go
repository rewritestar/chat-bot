package service

import (
	"chat-bot/src/app/ws/values"
	"chat-bot/src/common-service/chat"
	"chat-bot/src/core/ollama"
	"chat-bot/src/core/ws"

	"github.com/gin-gonic/gin"
)

type wsService struct {
	chatService  chat.ChatService
	ollamService ollama.OllamaService
}

func NewWsService(chatService chat.ChatService, ollamService ollama.OllamaService) WsService {
	return &wsService{
		chatService,
		ollamService,
	}
}

func (s *wsService) AddClient(ctx *gin.Context) error {
	conn, err := values.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return err
	}
	client := ws.NewClient(ws.GetHub(), conn, make(chan []byte, 256), ctx, s.chatService, s.ollamService)

	go client.WritePump()
	go client.ReadPump()
	return nil
}
