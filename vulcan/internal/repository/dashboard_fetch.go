package repository

import (
	"vulcan/internal/models"
)

func (repo *Repository) DashboardFetch() ([]models.Link, error) {
	var links []models.Link
	db := repo.DB

	err := db.
		Find(&links).
		Error

	return links, err
}
