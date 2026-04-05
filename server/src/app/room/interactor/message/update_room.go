package message

import "chat-bot/src/app/room/domain"

type UpdateRoom struct {
	Name        string `json:"name"`
	AiSystem    string `json:"aiSystem"`
	IsProactive *bool  `json:"isProactive"`
}

func (r *UpdateRoom) ToRoom(id, workerID uint) *domain.Room {
	result := domain.Room{}
	result.ID = id
	result.Name = r.Name
	result.AiSystem = r.AiSystem
	result.IsProactive = r.IsProactive
	result.CreatorID = workerID
	return &result
}
