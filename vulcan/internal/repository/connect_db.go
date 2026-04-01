package repository

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *Repository {
	log.Println("called")

	dsn := "user:password@tcp(127.0.0.1:3306)/wintermute?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect")
	}

	log.Println("db connected")

	return &Repository{DB: db}
}
