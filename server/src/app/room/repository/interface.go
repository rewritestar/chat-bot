package repository

import "chat-bot/src/app/room/domain"

type RoomRepository interface {
	FindAllRoom(uint) (*domain.RoomList, error)
	FindRoomByID(uint) (*domain.Room, error)
	SaveRoom(domain.Room) (*domain.Room, error)
	UpdateRoom(domain.Room) (*domain.Room, error)
}
