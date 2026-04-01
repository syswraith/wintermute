package repository

import (
	"time"

	"vulcan/internal/models"
)

func (repo *Repository) Create(longURL string, expiry string) (string, error) {
	var expiresAt *time.Time
	db := repo.DB

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
