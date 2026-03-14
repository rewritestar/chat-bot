package service

import "chat-bot/src/core/ws/internal/domain"

type ChatService interface {
	SaveChat(domain.Chat) (*domain.Chat, error)
}
