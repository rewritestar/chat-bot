package repository

import (
	"log"

	"chat-bot/src/app/room/domain"
	"chat-bot/src/initial/default_data"

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
	aiWorkerID := default_data.GetAIWorker().ID

	lastChatQuery := r.db.
		Model(&domain.Chat{}).
		Select("room_id, MAX(date_created) as last_chat_date").
		Group("room_id")

	readCountQuery := r.db.
		Model(&domain.Chat{}).
		Joins("JOIN room ON room.id = chat.room_id").
		Select("room_id, COUNT(*) AS unread_count").
		Where("chat.id > room.last_read_chat_id").
		Where("chat.creator_id = ?", aiWorkerID).
		Group("room_id")

	err := r.db.Model(&result).
		Select("room.*, COALESCE(rcq.unread_count, 0) as unread_count").
		Joins("LEFT JOIN (?) as lcq ON lcq.room_id = id", lastChatQuery).
		Joins("LEFT JOIN (?) as rcq ON rcq.room_id = id", readCountQuery).
		Where("creator_id = ? AND is_deleted = ?", workerID, false).
		Preload("Creator").
		Preload("ChatList", func(db *gorm.DB) *gorm.DB {
			return db.Order("date_created DESC")
		}).
		Order("lcq.last_chat_date IS NOT NULL, lcq.last_chat_date DESC").
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

func (r *roomRepository) UpdateLastReadChatID(id uint) error {
	var lastChatID *uint

	err := r.db.Model(&domain.Chat{}).
		Select("MAX(chat.id) AS last_chat_id").
		Joins("JOIN room ON chat.room_id = room.id").
		Where("room.id = ?", id).
		Where("chat.creator_id = ?", default_data.GetAIWorker().ID).
		Scan(&lastChatID).
		Error
	if err != nil {
		return err
	}

	var result domain.Room
	result.ID = id
	result.LastReadChatID = lastChatID

	if err = r.db.Updates(&result).Error; err != nil {
		return err
	}
	return nil
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
