package push

import (
	"chat-bot/src/common-service/push/domain"
	"chat-bot/src/common-service/push/repository"
)

type pushService struct {
	repo repository.PushRepository
}

func NewPushService(repo repository.PushRepository) PushService {
	return &pushService{
		repo,
	}
}

func (s *pushService) FindAllByUserID(userID uint) (*domain.PushSubscriptionList, error) {
	return s.repo.FindAllByUserID(userID)
}
