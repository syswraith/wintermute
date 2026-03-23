package models

import "time"

type Link struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	ShortURL string `gorm:"size:16,uniqueIndex"`
	LongURL  string `gorm:"type:text"`
	Expiry   *time.Time
}
