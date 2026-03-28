package chat

import "chat-bot/src/common-service/chat/domain"

type ChatService interface {
	SaveChat(domain.Chat) (*domain.Chat, error)
	UpdateRoom(domain.Room) (*domain.Room, error)

	FindHistoryByRoomID(uint) (*domain.Room, error)
	FindRoomByCreatorID(uint) (*domain.RoomList, error)
	FindAllRoomTickAiSchedule() (*domain.RoomList, error)
}
