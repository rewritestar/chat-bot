package domain

import "chat-bot/src/model"

type PushSubscription struct {
	model.PushSubscription
}

type RequestPush struct {
	UserID uint

	RoomID uint

	Title string

	Body string
}

type RequestMessage struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Url   string `json:"url"`
}
