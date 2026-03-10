package domain

import "chat-bot/src/model"

type Chat struct {
	model.Chat
}

type ChatList []*Chat
