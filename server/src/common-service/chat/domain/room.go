package domain

import "chat-bot/src/model"

type Room struct {
	model.Room

	ChatList ChatList
}

type RoomList []Room

func (r *RoomList) GetIDs() []uint {
	result := []uint{}

	for _, room := range *r {
		result = append(result, room.ID)
	}
	return result
}
