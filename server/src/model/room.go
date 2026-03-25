package model

import "time"

type Room struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"type:varchar(255)"`
	AiSystem    string    `gomr:"type:varchat(10000)"`
	CreatorID   uint      `gorm:"not null"`
	DateCreated time.Time `gorm:"autoCreateTime"`
	LastUpdated time.Time `gorm:"autoUpdateTime"`
}

func (*Room) TableName() string {
	return "room"
}
