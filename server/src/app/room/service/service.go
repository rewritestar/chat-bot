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

func (s *roomService) SaveRoom(reqData domain.Room) (*domain.Room, error) {
	return s.repo.SaveRoom(reqData)
}

func (s *roomService) UpdateRoom(reqData domain.Room) (*domain.Room, error) {
	return s.repo.UpdateRoom(reqData)
}

func (s *roomService) ShowRoom() (*domain.Room, error) {
	return s.repo.FindRoom()
}

func (s *roomService) SingleRoom(reqData domain.Room) (*domain.Room, error) {
	var room *domain.Room
	var err error
	room, err = s.repo.FindRoom()
	if err != nil {
		return nil, err
	}
	if room.ID == 0 {
		if _, err := s.repo.SaveRoom(reqData); err != nil {
			return nil, err
		}
		room, err = s.repo.FindRoom()
		if err != nil {
			return nil, err
		}
	}
	return room, nil
}
