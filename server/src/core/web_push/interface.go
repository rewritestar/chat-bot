package webpush

import "chat-bot/src/core/web_push/domain"

type WebPushService interface {
	Push(domain.RequestPush)
}
