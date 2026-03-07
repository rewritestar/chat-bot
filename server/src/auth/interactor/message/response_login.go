package message

import (
	"time"

	"chat-bot/src/auth/domain"
)

type ResponseLogin struct {
	Token string    `json:"token"`
	Exp   time.Time `json:"exp"`
}

func (r *ResponseLogin) Build(token *domain.Token) {
	r.Token = token.Token
	r.Exp = token.Exp
}
