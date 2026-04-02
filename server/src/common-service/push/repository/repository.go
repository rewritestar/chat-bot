package repository

import (
	"log"

	"chat-bot/src/common-service/push/domain"

	"gorm.io/gorm"
)

type PushRepository interface {
	FindAllByUserID(uint) (*domain.PushSubscriptionList, error)
}

type pushRepository struct {
	db *gorm.DB
}

func NewPushRepository(db *gorm.DB) PushRepository {
	r := &pushRepository{}
	r.db = db
	return r
}

func (r *pushRepository) FindAllByUserID(userID uint) (*domain.PushSubscriptionList, error) {
	result := domain.PushSubscriptionList{}

	err := r.db.Model(&result).
		Where("user_id = ?", userID).
		Order("date_created DESC").
		Find(&result).
		Error

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &result, nil
}
