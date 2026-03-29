package repository

import (
	"log"

	"chat-bot/src/app/room/domain"

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

func (r *roomRepository) FindAllRoom(workerID uint) (*domain.RoomList, error) {
	result := domain.RoomList{}

	subQuery := r.db.
		Model(&domain.Chat{}).
		Select("room_id, MAX(date_created) as last_chat_date").
		Group("room_id")

	err := r.db.Model(&result).
		Joins("LEFT JOIN (?) as last_chat ON last_chat.room_id = id", subQuery).
		Where("creator_id = ? AND is_deleted = ?", workerID, false).
		Preload("Creator").
		Preload("ChatList", func(db *gorm.DB) *gorm.DB {
			return db.Order("date_created DESC")
		}).
		Order("last_chat.last_chat_date IS NOT NULL, last_chat.last_chat_date DESC").
		Find(&result).
		Error

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &result, nil
}

func (r *roomRepository) FindRoomByID(roomID uint) (*domain.Room, error) {
	result := domain.Room{}
	result.ID = roomID

	err := r.db.Model(&result).
		Preload("Creator").
		Preload("ChatList", func(db *gorm.DB) *gorm.DB {
			return db.Order("date_created")
		}).
		First(&result).
		Error

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &result, nil
}

func (r *roomRepository) SaveRoom(room domain.Room) (*domain.Room, error) {
	if err := r.db.Create(&room).Error; err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &room, nil
}

func (r *roomRepository) UpdateRoom(room domain.Room) (*domain.Room, error) {
	if err := r.db.Updates(&room).Error; err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &room, nil
}

func (r *roomRepository) SoftDeleteRoom(id uint) error {
	err := r.db.Model(&domain.Room{}).
		Where("id = ?", id).
		UpdateColumns(map[string]interface{}{
			"is_deleted": true,
		}).
		Error

	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
