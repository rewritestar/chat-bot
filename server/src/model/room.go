package model

import "time"

type Room struct {
	ID                  uint       `gorm:"primaryKey;autoIncrement"`
	Name                string     `gorm:"type:varchar(255)"`
	AiSystem            string     `gorm:"type:varchar(10000)"`
	Memory              string     `gorm:"text"`
	MemoryPendingCount  *uint      `gorm:"not null;default:0"`
	AiProactiveNextDate *time.Time `gorm:""`
	LastReadChatID      *uint      `gorm:""`
	IsProactive         *bool      `gorm:"not null;default:true"`
	CreatorID           uint       `gorm:"not null"`
	DateCreated         time.Time  `gorm:"autoCreateTime"`
	LastUpdated         time.Time  `gorm:"autoUpdateTime"`
	MemoryUpdated       time.Time  `gorm:""`
}

func (*Room) TableName() string {
	return "room"
}
