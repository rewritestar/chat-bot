package domain

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"chat-bot/src/auth/values"
	"chat-bot/src/model"
)

type Worker struct {
	model.Worker
}

func (w *Worker) GenerateToken() (*Token, error) {
	exp := jwt.NewNumericDate(time.Now().Add(10 * time.Second))
	t := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"workerId": w.ID,
			"exp":      exp,
		},
	)
	token, err := t.SignedString([]byte(os.Getenv(values.EnvJwtSecret)))
	if err != nil {
		return nil, err
	}
	result := &Token{
		Token: token,
		Exp:   exp.Time,
	}
	return result, nil
}
