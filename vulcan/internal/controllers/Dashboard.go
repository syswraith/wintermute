package controllers

import (
	"log"
	"net/http"

	"vulcan/internal/repository"

	"github.com/gin-gonic/gin"
)

func (c *Control) Dashboard(context *gin.Context) {
	links, err := repository.DashboardFetch(c.DB)
	if err != nil {
		log.Println("error fetching dashboard links:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "dashboard links fetching failed"})
		return
	}

	context.JSON(http.StatusOK, links)
}

