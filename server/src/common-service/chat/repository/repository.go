package repository

import (
	"chat-bot/src/common-service/chat/domain"
	"log"

	"gorm.io/gorm"
)

type ChatRepository interface {
	SaveChat(domain.Chat) (*domain.Chat, error)
}

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
