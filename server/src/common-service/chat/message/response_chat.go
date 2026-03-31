package message

import "chat-bot/src/common-service/chat/domain"

type ResponseChat struct {
	ID          uint   `json:"id"`
	Content     string `json:"content"`
	CreatorID   uint   `json:"creatorId"`
	DateCreated string `json:"dateCreated"`
}

func (r *ResponseChat) Build(chat *domain.Chat) {
	r.ID = chat.ID
	r.Content = chat.Content
	r.CreatorID = chat.CreatorID
	r.DateCreated = chat.DateCreated.Format("06.01.02 15:04pm")
}
