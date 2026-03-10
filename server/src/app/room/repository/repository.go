package repository

import (
	"chat-bot/src/app/room/domain"
	"log"

	"gorm.io/gorm"
)

type roomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) RoomRepository {
	r := &roomRepository{}
	r.db = db
	return r
}

func (r *roomRepository) SaveRoom(room domain.Room) (*domain.Room, error) {
	if err := r.db.Create(&room).Error; err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &room, nil
}

func (r *roomRepository) FindRoom() (*domain.Room, error) {
	result := domain.Room{}

	err := r.db.Model(&result).
		Preload("Creator").
		Preload("ChatList", func(db *gorm.DB) *gorm.DB {
			return db.Order("date_created")
		}).
		Limit(1).
		Find(&result).
		Error

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &result, nil
}

func (r *roomRepository) SaveChat(chat domain.Chat) (*domain.Chat, error) {
	if err := r.db.Create(&chat).Error; err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &chat, nil
}
