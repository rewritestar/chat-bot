package repository

import (
	"log"

	"chat-bot/src/core/ws/internal/domain"

	"gorm.io/gorm"
)

type chatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) ChatRepository {
	r := &chatRepository{}
	r.db = db
	return r
}

func (r *chatRepository) SaveChat(chat domain.Chat) (*domain.Chat, error) {
	if err := r.db.Create(&chat).Error; err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &chat, nil
}
