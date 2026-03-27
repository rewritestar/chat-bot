package domain

import "chat-bot/src/model"

type Room struct {
	model.Room

	Creator  Worker   `gorm:"foreignKey:creatorID"`
	ChatList ChatList `gorm:"foreignKey:RoomID"`
}

type RoomList []Room
