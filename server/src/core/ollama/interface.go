package ollama

import "chat-bot/src/core/ollama/domain"

type OllamaService interface {
	Chat(string) *domain.ResponseBody
}
