package model

import "time"

type Worker struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Email       string    `gorm:"type:varchar(100);not null"`
	Password    string    `gorm:"type:varchar(255);not null"`
	DateCreated time.Time `gorm:"autoCreateTime"`
	LastUpdated time.Time `gorm:"autoUpdateTime"`
}

func (*Worker) TableName() string {
	return "worker"
}
