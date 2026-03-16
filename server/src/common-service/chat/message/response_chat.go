package message

import "chat-bot/src/common-service/chat/domain"

type ResponseChat struct {
	ID        uint   `json:"id"`
	Content   string `json:"content"`
	CreatorID uint   `json:"creatorId"`
}

func (r *ResponseChat) Build(chat *domain.Chat) {
	r.ID = chat.ID
	r.Content = chat.Content
	r.CreatorID = chat.CreatorID
}
