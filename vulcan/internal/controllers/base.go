package controllers

import "gorm.io/gorm"

type Control struct {
	DB *gorm.DB
}

func Init(db *gorm.DB) *Control {
	return &Control{DB: db}
}
