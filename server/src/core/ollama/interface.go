package ollama

import (
	chatDomain "chat-bot/src/common-service/chat/domain"
	"chat-bot/src/core/ollama/domain"
)

type OllamaService interface {
	Chat(domain.RequestChat, chatDomain.Room) *domain.ResponseBody
}
