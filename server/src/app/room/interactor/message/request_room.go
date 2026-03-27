package message

import "chat-bot/src/app/room/domain"

type RequestRoom struct {
	Name     string `json:"name"`
	AiSystem string `json:"aiSystem"`
}

func (r *RequestRoom) ToRoom(workerID uint) *domain.Room {
	result := domain.Room{}
	result.Name = r.Name
	result.AiSystem = r.AiSystem
	result.CreatorID = workerID
	return &result
}
