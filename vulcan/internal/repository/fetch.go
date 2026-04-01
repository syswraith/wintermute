package repository

import (
	"errors"
	"time"

	"vulcan/internal/models"
)

func (repo *Repository) Fetch(shortURL string) (string, error) {
	var link models.Link
	db := repo.DB

	err := db.
		Where("short_url = ?", shortURL).
		First(&link).
		Error

	if link.Expiry != nil && time.Now().After(*link.Expiry) {
		err = errors.New("link expired")
	}

	return link.LongURL, err
}
