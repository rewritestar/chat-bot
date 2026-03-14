package repository

import "chat-bot/src/core/ws/internal/domain"

type ChatRepository interface {
	SaveChat(domain.Chat) (*domain.Chat, error)
}
