package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	core_values "chat-bot/src/core/values"
)

var db *gorm.DB

func NewMariaDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/chatbot?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv(core_values.EnvDBUser), os.Getenv(core_values.EnvDBPW), os.Getenv(core_values.EnvDBHost))
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = gormDB
}

func GetMariaDB() *gorm.DB {
	return db
}
