package message

import "chat-bot/src/app/room/domain"

type ResponseRoom struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatorID uint   `json:"creatorId"`

	ChatList []responseChat `json:"chatList"`
}

type responseChat struct {
	ID        uint   `json:"id"`
	Content   string `json:"content"`
	CreatorID uint   `json:"creatorId"`
}

func (r *ResponseRoom) Build(room domain.Room) {
	r.ID = room.ID
	r.Name = room.Name
	r.CreatorID = room.CreatorID

	for _, chat := range room.ChatList {
		responseChat := responseChat{}
		responseChat.ID = chat.ID
		responseChat.Content = chat.Content
		responseChat.CreatorID = chat.CreatorID

		r.ChatList = append(r.ChatList, responseChat)
	}
}
