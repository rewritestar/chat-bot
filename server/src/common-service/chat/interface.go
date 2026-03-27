package chat

import "chat-bot/src/common-service/chat/domain"

type ChatService interface {
	SaveChat(domain.Chat) (*domain.Chat, error)
	FindHistoryByRoomID(uint) (*domain.Room, error)
	FindRoomByCreatorID(uint) (*domain.RoomList, error)
}
