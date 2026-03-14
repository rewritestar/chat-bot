package service

import (
	"chat-bot/src/core/ws/internal/domain"
	"chat-bot/src/core/ws/internal/repository"
)

type chatService struct {
	repo repository.ChatRepository
}

func NewChatService(repo repository.ChatRepository) ChatService {
	return &chatService{
		repo,
	}
}

func (s *chatService) SaveChat(reqData domain.Chat) (*domain.Chat, error) {
	return s.repo.SaveChat(reqData)
}
