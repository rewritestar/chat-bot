package service

import (
	"chat-bot/src/app/push/domain"
	"chat-bot/src/app/push/repository"
)

type pushService struct {
	repo repository.PushRepository
}

func NewService(repo repository.PushRepository) PushService {
	return &pushService{
		repo,
	}
}

func (s *pushService) Save(reqData domain.PushSubscription) (*domain.PushSubscription, error) {
	return s.repo.UpsertPush(reqData)
}
