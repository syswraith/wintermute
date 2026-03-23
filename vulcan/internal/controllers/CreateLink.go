package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"vulcan/internal/models"
	"vulcan/internal/repository"
)

func (c *Control) CreateLink(context *gin.Context) {
	var json models.PostJSON

	if err := context.ShouldBindJSON(&json); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	shortURL, err := repository.Create(json.LongURL, json.Expiry, c.DB)
	if err != nil {
		log.Println("error generating link:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "link generation failed"})
		return
	}

	shortURL = "http://localhost:31337/" + shortURL

	context.JSON(http.StatusOK, gin.H{"shortURL": shortURL})
}

