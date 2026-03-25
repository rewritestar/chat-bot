package domain

import "chat-bot/src/model"

type Chat struct {
	model.Chat
}

type ChatList []Chat

func (c ChatList) Len() int {
	return len(c)
}

func (c ChatList) Less(i, j int) bool {
	return c[i].DateCreated.Before(c[j].DateCreated)
}

func (c ChatList) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
