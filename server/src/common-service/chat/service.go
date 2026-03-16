package chat

import (
	"chat-bot/src/common-service/chat/domain"
	"chat-bot/src/common-service/chat/repository"
)

type chatService struct {
	repo repository.ChatRepository
}

func NewChatService(repo repository.ChatRepository) ChatService {
	return &chatService{
		repo,
	}
}

func (s *chatService) SaveChat(chat domain.Chat) (*domain.Chat, error) {
	return s.repo.SaveChat(chat)
}
