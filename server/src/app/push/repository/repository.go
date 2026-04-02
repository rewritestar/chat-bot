package repository

import (
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"chat-bot/src/app/push/domain"
)

type pushRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) PushRepository {
	r := &pushRepository{}
	r.db = db
	return r
}

func (r *pushRepository) UpsertPush(push domain.PushSubscription) (*domain.PushSubscription, error) {
	err := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "endpoint"}},
		DoUpdates: clause.AssignmentColumns([]string{"user_id", "p256dh", "auth", "last_updated"}),
	}).Create(&push).
		Error

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &push, nil
}
