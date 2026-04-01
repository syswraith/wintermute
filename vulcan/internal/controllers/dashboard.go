package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Dashboard(context *gin.Context) {
	links, err := h.Repo.DashboardFetch()
	if err != nil {
		log.Println("error fetching dashboard links:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "dashboard links fetching failed"})
		return
	}

	context.JSON(http.StatusOK, links)
}
