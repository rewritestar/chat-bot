package service

import "chat-bot/src/auth/domain"

type AuthService interface {
	Signin(domain.Worker) (*domain.Worker, error)
	Login(domain.Worker) (*domain.Token, error)
}
