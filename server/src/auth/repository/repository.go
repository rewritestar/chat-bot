package repository

import (
	"chat-bot/src/auth/domain"
	"log"

	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	r := &authRepository{}
	r.db = db
	return r
}

func (r *authRepository) SaveWorker(worker domain.Worker) (*domain.Worker, error) {
	if err := r.db.Create(&worker).Error; err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &worker, nil
}

func (r *authRepository) FindWorkerByEmail(email string) (*domain.Worker, error) {
	result := domain.Worker{}

	err := r.db.Where("email = ?", email).
		First(&result).
		Error

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &result, nil
}
