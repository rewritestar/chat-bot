package domain

import "chat-bot/src/model"

type Room struct {
	model.Room

	ChatList ChatList
}
