package message

import "chat-bot/src/app/push/domain"

type RequestPush struct {
	Endpoint string `json:"endpoint" binding:"required"`
	P256dh   string `json:"p256dh" binding:"required"`
	Auth     string `json:"auth" binding:"required"`
}

func (r *RequestPush) ToPushSub(userID uint) *domain.PushSubscription {
	result := domain.PushSubscription{}
	result.UserID = userID
	result.Endpoint = r.Endpoint
	result.P256dh = r.P256dh
	result.Auth = r.Auth
	return &result
}
