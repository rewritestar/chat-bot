package service

import (
	"chat-bot/src/app/ws/values"
	"chat-bot/src/common-service/chat"
	"chat-bot/src/core/ollama"
	webpush "chat-bot/src/core/web_push"
	"chat-bot/src/core/ws"

	"github.com/gin-gonic/gin"
)

type wsService struct {
	chatService  chat.ChatService
	ollamService ollama.OllamaService
	pushService  webpush.WebPushService
}

func NewWsService(chatService chat.ChatService, ollamService ollama.OllamaService, pushService webpush.WebPushService,
) WsService {
	return &wsService{
		chatService,
		ollamService,
		pushService,
	}
}

func (s *wsService) AddClient(ctx *gin.Context) error {
	conn, err := values.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return err
	}
	client := ws.NewClient(ws.GetHub(), conn, make(chan []byte, 256), s.chatService, s.ollamService, s.pushService)

	go client.WritePump()
	go client.ReadPump()
	return nil
}
