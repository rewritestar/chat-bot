package service

import "chat-bot/src/app/room/domain"

type RoomService interface {
	SaveRoom(domain.Room) (*domain.Room, error)
	ShowRoom() (*domain.Room, error)
	SingleRoom(domain.Room) (*domain.Room, error)
}
