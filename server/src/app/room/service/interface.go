package service

import "chat-bot/src/app/room/domain"

type RoomService interface {
	IndexRoom(uint) (*domain.RoomList, error)
	ShowRoom(uint) (*domain.Room, error)
	SaveRoom(domain.Room) (*domain.Room, error)
	UpdateRoom(domain.Room) (*domain.Room, error)
}
