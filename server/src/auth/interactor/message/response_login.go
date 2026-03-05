package message

import "chat-bot/src/auth/domain"

type ResponseLogin struct {
	WorkerID uint   `json:"workerId"`
	Email    string `json:"email"`
	Token    string // 구현 필요
}

func (r *ResponseLogin) Build(worker domain.Worker) {
	r.WorkerID = worker.ID
	r.Email = worker.Email
}
