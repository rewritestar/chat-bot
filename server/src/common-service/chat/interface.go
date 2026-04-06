package chat

import "chat-bot/src/common-service/chat/domain"

type ChatService interface {
	SaveChat(domain.Chat) (*domain.Chat, error)
	UpdateRoom(domain.Room) (*domain.Room, error)
	UpdateMemoryPendingCount(uint) (*domain.Room, error)

	FindHistoryByRoomID(uint, int) (*domain.Room, error)
	FindRoomByCreatorID(uint) (*domain.RoomList, error)
	FindRoomByID(uint) (*domain.Room, error)
	FindAllRoomTickAiSchedule() (*domain.RoomList, error)
}
