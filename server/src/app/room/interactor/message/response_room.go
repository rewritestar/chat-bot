package message

import "chat-bot/src/app/room/domain"

type ResponseRoom struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatorID uint   `json:"creatorId"`

	ChatList []ResponseChat `json:"chatList"`
}

type ResponseChat struct {
	ID        uint   `json:"id"`
	Content   string `json:"content"`
	CreatorID uint   `json:"creatorId"`
}

func (r *ResponseRoom) Build(room domain.Room) {
	r.ID = room.ID
	r.Name = room.Name
	r.CreatorID = room.CreatorID

	for _, chat := range room.ChatList {
		responseChat := ResponseChat{}
		responseChat.Build(chat)

		r.ChatList = append(r.ChatList, responseChat)
	}
}

func (r *ResponseChat) Build(chat *domain.Chat) {
	r.ID = chat.ID
	r.Content = chat.Content
	r.CreatorID = chat.CreatorID
}
