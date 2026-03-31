package message

import "chat-bot/src/app/room/domain"

type ResponseRoomList struct {
	List []ResponseRoom `json:"list"`
}

type ResponseRoom struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	AiSystem       string `json:"aiSystem"`
	LastReadChatID *uint  `json:"lastReadChatID"`
	UnreadCount    uint   `json:"unreadCount"`
	CreatorID      uint   `json:"creatorId"`

	ChatList []responseChat `json:"chatList"`
}

type responseChat struct {
	ID          uint   `json:"id"`
	Content     string `json:"content"`
	CreatorID   uint   `json:"creatorId"`
	DateCreated string `json:"dateCreated"`
}

func (r *ResponseRoomList) Build(roomList domain.RoomList) {
	for _, room := range roomList {
		respRoom := ResponseRoom{}
		respRoom.build(room)
		r.List = append(r.List, respRoom)
	}
}

func (r *ResponseRoom) build(room domain.Room) {
	r.ID = room.ID
	r.Name = room.Name
	r.AiSystem = room.AiSystem
	r.LastReadChatID = room.LastReadChatID
	r.UnreadCount = room.UnreadCount
	r.CreatorID = room.CreatorID

	if len(room.ChatList) > 0 {
		responseChat := responseChat{}
		responseChat.Build(room.ChatList[0])
		r.ChatList = append(r.ChatList, responseChat)
	}
}

func (r *ResponseRoom) Build(room domain.Room) {
	r.ID = room.ID
	r.Name = room.Name
	r.AiSystem = room.AiSystem
	r.LastReadChatID = room.LastReadChatID
	r.UnreadCount = room.UnreadCount
	r.CreatorID = room.CreatorID

	for _, chat := range room.ChatList {
		responseChat := responseChat{}
		responseChat.Build(chat)
		r.ChatList = append(r.ChatList, responseChat)
	}
}

func (r *responseChat) Build(chat *domain.Chat) {
	r.ID = chat.ID
	r.Content = chat.Content
	r.CreatorID = chat.CreatorID
	r.DateCreated = chat.DateCreated.Format("06.01.02 15:04pm")
}
