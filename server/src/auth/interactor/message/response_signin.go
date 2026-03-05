package message

import "chat-bot/src/auth/domain"

type ResponseSignin struct {
	WorkerID uint   `json:"workerId"`
	Email    string `json:"email"`
}

func (r *ResponseSignin) Build(worker domain.Worker) {
	r.WorkerID = worker.ID
	r.Email = worker.Email
}
