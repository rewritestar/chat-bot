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

func (s *chatService) UpdateRoom(room domain.Room) (*domain.Room, error) {
	return s.repo.UpdateRoom(room)
}

func (s *chatService) UpdateMemoryPendingCount(roomID uint) (*domain.Room, error) {
	if err := s.repo.UpdateMemoryPendingCount(roomID); err != nil {
		return nil, err
	}
	return s.FindRoomByID(roomID)
}

func (s *chatService) FindHistoryByRoomID(roomID uint, limit int) (*domain.Room, error) {
	return s.repo.FindHistoryByRoomID(roomID, limit)
}

func (s *chatService) FindRoomByCreatorID(creatorID uint) (*domain.RoomList, error) {
	return s.repo.FindRoomByCreatorID(creatorID)
}

func (s *chatService) FindRoomByID(roomID uint) (*domain.Room, error) {
	return s.repo.FindRoomByID(roomID)
}

func (s *chatService) FindAllRoomTickAiSchedule() (*domain.RoomList, error) {
	return s.repo.FindAllRoomTickAiSchedule()
}
