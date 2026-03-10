package model

import "time"

type Chat struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Content     string    `gorm:"type:text"`
	RoomID      uint      `gorm:"not null"`
	CreatorID   uint      `gorm:"not null"`
	DateCreated time.Time `gorm:"autoCreateTime"`
	LastUpdated time.Time `gorm:"autoUpdateTime"`
}

func (*Chat) TableName() string {
	return "chat"
}
