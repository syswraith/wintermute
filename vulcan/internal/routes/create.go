package routes

import (
	"vulcan/internal/controllers"
	"vulcan/internal/middleware"

	"github.com/gin-gonic/gin"
)

func CreateRoute(router *gin.Engine, h *controllers.Handler) {
	router.POST("/create", middleware.RequireAuth(), h.CreateLink)
}
