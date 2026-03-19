package domain

import (
	"os"
	"time"

	core_values "chat-bot/src/core/values"
	"chat-bot/src/model"

	"github.com/golang-jwt/jwt/v5"
)

type Worker struct {
	model.Worker
}

func (w *Worker) GenerateToken() (*Token, error) {
	exp := jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
	t := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"workerId": w.ID,
			"exp":      exp,
		},
	)
	token, err := t.SignedString([]byte(os.Getenv(core_values.EnvJwtSecret)))
	if err != nil {
		return nil, err
	}
	result := &Token{
		Token:    token,
		Exp:      exp.Time,
		WorkerID: w.ID,
	}
	return result, nil
}
