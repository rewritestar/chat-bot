package repository

import "chat-bot/src/app/room/domain"

type RoomRepository interface {
	SaveRoom(domain.Room) (*domain.Room, error)
	UpdateRoom(domain.Room) (*domain.Room, error)
	FindRoom() (*domain.Room, error)
}
