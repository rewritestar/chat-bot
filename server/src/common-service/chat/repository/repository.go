package repository

import (
	"log"
	"sort"
	"time"

	"chat-bot/src/common-service/chat/domain"
	"chat-bot/src/common-service/chat/values"

	"gorm.io/gorm"
)

type ChatRepository interface {
	SaveChat(domain.Chat) (*domain.Chat, error)
	UpdateRoom(domain.Room) (*domain.Room, error)

	FindHistoryByRoomID(uint) (*domain.Room, error)
	FindRoomByCreatorID(uint) (*domain.RoomList, error)
	FindAllRoomTickAiSchedule() (*domain.RoomList, error)
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

func (r *chatRepository) UpdateRoom(room domain.Room) (*domain.Room, error) {
	if err := r.db.Updates(&room).Error; err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &room, nil
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

func (r *chatRepository) FindRoomByCreatorID(creatorID uint) (*domain.RoomList, error) {
	result := domain.RoomList{}

	err := r.db.Model(&result).
		Where("creator_id = ?", creatorID).
		Find(&result).
		Error

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &result, nil
}

func (r *chatRepository) FindAllRoomTickAiSchedule() (*domain.RoomList, error) {
	result := domain.RoomList{}

	err := r.db.Model(&result).
		Where(`ai_proactive_next_date <= ? OR ai_proactive_next_date IS NULL`, time.Now()).
		Where("is_deleted = ?", false).
		Preload("ChatList", func(db *gorm.DB) *gorm.DB {
			return db.Order("date_created DESC")
		}).
		Find(&result).
		Error

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &result, nil
}
