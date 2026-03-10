package repository

import "chat-bot/src/app/auth/domain"

type AuthRepository interface {
	SaveWorker(domain.Worker) (*domain.Worker, error)
	FindWorkerByEmail(string) (*domain.Worker, error)
}
