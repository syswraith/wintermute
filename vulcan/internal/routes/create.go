package routes

import (
	"vulcan/internal/controllers"

	"github.com/gin-gonic/gin"
)

func CreateRoute(router *gin.Engine, h *controllers.Handler) {
	router.POST("/create", h.CreateLink)
}

