package repository

import (
	"log"
	"sort"

	"chat-bot/src/common-service/chat/domain"
	"chat-bot/src/common-service/chat/values"

	"gorm.io/gorm"
)

type ChatRepository interface {
	SaveChat(domain.Chat) (*domain.Chat, error)
	FindHistoryByRoomID(uint) (*domain.Room, error)
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

func (r *chatRepository) FindHistoryByRoomID(roomID uint) (*domain.Room, error) {
	result := domain.Room{}
	result.ID = roomID

	if err := r.db.Model(&result).First(&result).Error; err != nil {
		log.Println(err.Error())
		return nil, err
	}

	err := r.db.Model(&result.ChatList).
		Where("room_id = ?", roomID).
		Order("date_created DESC").
		Limit(values.HistoryLimit).
		Find(&result.ChatList).
		Error

	sort.Sort(result.ChatList)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &result, nil
}
