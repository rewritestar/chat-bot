package service

import "chat-bot/src/app/push/domain"

type PushService interface {
	Save(domain.PushSubscription) (*domain.PushSubscription, error)
}
