package message

import "chat-bot/src/app/room/domain"

type RequestRoom struct {
	Name string `json:"name"`
}

func (r *RequestRoom) ToRoom(workerID uint) *domain.Room {
	result := domain.Room{}
	result.Name = r.Name
	result.CreatorID = workerID
	return &result
}
