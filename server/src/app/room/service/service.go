package service

import (
	"chat-bot/src/app/room/domain"
	"chat-bot/src/app/room/repository"
)

type roomService struct {
	repo repository.RoomRepository
}

func NewRoomService(repo repository.RoomRepository) RoomService {
	return &roomService{
		repo,
	}
}

func (s *roomService) IndexRoom(workerID uint) (*domain.RoomList, error) {
	return s.repo.FindAllRoom(workerID)
}

func (s *roomService) ShowRoom(roomID uint) (*domain.Room, error) {
	return s.repo.FindRoomByID(roomID)
}

func (s *roomService) SaveRoom(reqData domain.Room) (*domain.Room, error) {
	savedRoom, err := s.repo.SaveRoom(reqData)
	if err != nil {
		return nil, err
	}
	return savedRoom, nil
}

func (s *roomService) UpdateRoom(reqData domain.Room) (*domain.Room, error) {
	return s.repo.UpdateRoom(reqData)
}

func (s *roomService) SoftDeleteRoom(id uint) error {
	return s.repo.SoftDeleteRoom(id)
}
