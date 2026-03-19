package ws

import (
	"net/http"

	"chat-bot/src/app/ws/interactor"
	"chat-bot/src/app/ws/service"
	"chat-bot/src/common-service/chat"
	chatRepository "chat-bot/src/common-service/chat/repository"
	"chat-bot/src/core/ollama"
	"chat-bot/src/initial/database"

	"github.com/gin-gonic/gin"
)

func Main(r *gin.RouterGroup) {
	chatRepo := chatRepository.NewChatRepository(database.GetMariaDB())
	chatService := chat.NewChatService(chatRepo)
	ollamService := ollama.NewOllamaService()
	service := service.NewWsService(chatService, ollamService)

	r.GET("", func(ctx *gin.Context) {
		if err := service.AddClient(ctx); err != nil {
			interactor.ErrorPresenter(ctx, http.StatusInternalServerError, err)
			return
		}
	})
}
