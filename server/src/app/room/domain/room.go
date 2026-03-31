package domain

import "chat-bot/src/model"

type Room struct {
	model.Room

	UnreadCount uint     `gorm:"->"`
	Creator     Worker   `gorm:"foreignKey:creatorID"`
	ChatList    ChatList `gorm:"foreignKey:RoomID"`
}

type RoomList []Room
