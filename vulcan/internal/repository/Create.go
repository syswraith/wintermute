package repository

import (
	"time"

	"gorm.io/gorm"

	"vulcan/internal/models"
)

func Create(longURL string, expiry string, db *gorm.DB) (string, error) {
	var expiresAt *time.Time

	if expiry != "" {
		t, err := time.ParseInLocation("2006-01-02T15:04", expiry, time.Local)
		if err != nil {
			return "", err
		}
		expiresAt = &t
	}

	l := models.Link{
		LongURL: longURL,
		Expiry:  expiresAt,
	}

	if err := db.Create(&l).Error; err != nil {
		return "", err
	}

	shortURL := ShorturlGenerator(l.ID)

	if err := db.Model(&l).Update("short_url", shortURL).Error; err != nil {
		return "", err
	}

	return shortURL, nil
}
