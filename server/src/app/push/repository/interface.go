package repository

import "chat-bot/src/app/push/domain"

type PushRepository interface {
	UpsertPush(domain.PushSubscription) (*domain.PushSubscription, error)
}
