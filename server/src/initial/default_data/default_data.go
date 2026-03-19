package default_data

import (
	"log"

	"chat-bot/src/app/auth/domain"

	"gorm.io/gorm"
)

var aiWorker *domain.Worker

func SetDefaultData(db *gorm.DB) {
	var worker *domain.Worker
	if err := db.First(&worker).Error; err != nil {
		log.Fatalln(err)
	}
	aiWorker = worker
}

func GetAIWorker() *domain.Worker {
	return aiWorker
}
