package repository

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"vulcan/internal/models"
)

func Fetch(shortURL string, db *gorm.DB) (string, error) {
	var link models.Link
	err := db.
		Where("short_url = ?", shortURL).
		First(&link).
		Error

	if link.Expiry != nil && time.Now().After(*link.Expiry) {
		err = errors.New("link expired")
	}

	return link.LongURL, err
}
