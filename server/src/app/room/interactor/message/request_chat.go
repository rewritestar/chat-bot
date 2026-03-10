package message

import "chat-bot/src/app/room/domain"

type RequestChat struct {
	RoomID  uint   `json:"roomId"`
	Content string `json:"content"`
}

func (r *RequestChat) ToChat(workerID uint) *domain.Chat {
	result := domain.Chat{}
	result.RoomID = r.RoomID
	result.Content = r.Content
	result.CreatorID = workerID
	return &result
}
