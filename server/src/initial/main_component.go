package initial

import (
	"log"

	"chat-bot/src/common-service/chat"
	"chat-bot/src/common-service/chat/repository"
	"chat-bot/src/core/ollama"
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

	ollamaSvc := ollama.NewOllamaService()
	chatRepo := repository.NewChatRepository(database.GetMariaDB())
	chatSvc := chat.NewChatService(chatRepo)
	hub := ws.GetHub()

	aiScheduler := aischeduler.NewAiSchduler(ollamaSvc, chatSvc, hub)
	aiScheduler.StartAIScheduler()
}
