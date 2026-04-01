package controllers

import (
	"vulcan/internal/repository"
)

type Handler struct {
	Repo *repository.Repository
}

func Init(repo *repository.Repository) *Handler {
	return &Handler{Repo: repo}
}
