package message

import (
	"time"

	"chat-bot/src/app/auth/domain"
)

type ResponseLogin struct {
	Token    string    `json:"token"`
	Exp      time.Time `json:"exp"`
	WorkerID uint      `json:"workerId"`
}

func (r *ResponseLogin) Build(token *domain.Token) {
	r.Token = token.Token
	r.Exp = token.Exp
	r.WorkerID = token.WorkerID
}
