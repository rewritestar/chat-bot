package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func NewMariaDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/chatbot?charset=utf8mb4&parseTime=True&loc=Local"
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = gormDB
}

func GetMariaDB() *gorm.DB {
	return db
}
