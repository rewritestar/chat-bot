package repository

import (
	"log"
	"sort"
	"time"

	"chat-bot/src/common-service/chat/domain"

	"gorm.io/gorm"
)

type ChatRepository interface {
	SaveChat(domain.Chat) (*domain.Chat, error)
	UpdateRoom(domain.Room) (*domain.Room, error)
	UpdateMemoryPendingCount(uint) error

	FindHistoryByRoomID(uint, int) (*domain.Room, error)
	FindRoomByCreatorID(uint) (*domain.RoomList, error)
	FindRoomByID(uint) (*domain.Room, error)
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

func (r *chatRepository) UpdateMemoryPendingCount(roomID uint) error {
	err := r.db.Model(&domain.Room{}).
		Where("room.id = ?", roomID).
		Update("memory_pending_count", gorm.Expr("memory_pending_count + ?", 1)).
		Error

	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return nil
}

func (r *chatRepository) FindHistoryByRoomID(roomID uint, limit int) (*domain.Room, error) {
	result := domain.Room{}
	result.ID = roomID

	if err := r.db.Model(&result).First(&result).Error; err != nil {
		log.Println(err.Error())
		return nil, err
	}

	err := r.db.Model(&result.ChatList).
		Where("room_id = ?", roomID).
		Order("date_created DESC").
		Limit(limit).
		Find(&result.ChatList).
		Error

	sort.Sort(result.ChatList)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &result, nil
}

func (r *chatRepository) FindRoomByID(roomID uint) (*domain.Room, error) {
	result := domain.Room{}
	result.ID = roomID

	err := r.db.Model(&result).
		First(&result).
		Error

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
		Where("is_proactive = ?", true).
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
