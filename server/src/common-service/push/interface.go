package push

import "chat-bot/src/common-service/push/domain"

type PushService interface {
	FindAllByUserID(uint) (*domain.PushSubscriptionList, error)
}
