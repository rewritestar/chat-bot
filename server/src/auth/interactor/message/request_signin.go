package message

import "chat-bot/src/auth/domain"

type RequestSignin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r *RequestSignin) ToWorker() *domain.Worker {
	result := domain.Worker{}
	result.Email = r.Email
	result.Password = r.Password
	return &result
}
