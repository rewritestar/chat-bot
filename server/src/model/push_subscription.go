package model

import "time"

type PushSubscription struct {
	UserID      uint      `gorm:"not null"`
	Endpoint    string    `gorm:"type:text;not null"`
	P256dh      string    `gorm:"type:text;not nul"`
	Auth        string    `gorm:"type:text;not nul"`
	DateCreated time.Time `gorm:"autoCreateTime"`
	LastUpdated time.Time `gorm:"autoUpdateTime"`
}

func (*PushSubscription) TableName() string {
	return "push_subscription"
}
