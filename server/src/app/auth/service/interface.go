package service

import "chat-bot/src/app/auth/domain"

type AuthService interface {
	Signin(domain.Worker) (*domain.Worker, error)
	Login(domain.Worker) (*domain.Token, error)
}
