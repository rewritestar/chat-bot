package domain

type Message struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

type AuthMessage struct {
	Token string `json:"token"`
}

type JoinMessage struct {
	RoomID uint `json:"roomId"`
}

type ChatMessage struct {
	RoomID  uint   `json:"roomId"`
	Content string `json:"content"`
}
