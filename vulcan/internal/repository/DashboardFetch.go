package repository

import (
	"gorm.io/gorm"

	"vulcan/internal/models"
)

func DashboardFetch(db *gorm.DB) ([]models.Link, error) {
	var links []models.Link
	err := db.
		Find(&links).
		Error

	return links, err
}
