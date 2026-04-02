package initial

import (
	"log"

	"chat-bot/src/common-service/chat"
	"chat-bot/src/common-service/chat/repository"
	"chat-bot/src/common-service/push"
	pushRepostiry "chat-bot/src/common-service/push/repository"
	"chat-bot/src/core/ollama"
	webpush "chat-bot/src/core/web_push"
	"chat-bot/src/core/ws"
	aischeduler "chat-bot/src/initial/ai_scheduler"
	"chat-bot/src/initial/database"
	"chat-bot/src/initial/default_data"

	"github.com/joho/godotenv"
)

func Main() {
	database.NewMariaDB()

	ws.NewHub()
	go ws.GetHub().Run()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	default_data.SetDefaultData(database.GetMariaDB())

	hub := ws.GetHub()
	ollamaSvc := ollama.NewOllamaService()
	chatRepo := repository.NewChatRepository(database.GetMariaDB())
	chatSvc := chat.NewChatService(chatRepo)
	pushRepo := pushRepostiry.NewPushRepository(database.GetMariaDB())
	pushSvc := push.NewPushService(pushRepo)
	webpushSvc := webpush.NewWebPush(pushSvc)

	aiScheduler := aischeduler.NewAiSchduler(hub, ollamaSvc, chatSvc, webpushSvc)
	aiScheduler.StartAIScheduler()
}
