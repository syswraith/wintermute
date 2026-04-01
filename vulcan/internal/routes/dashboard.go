package routes

import (
	"vulcan/internal/controllers"

	"github.com/gin-gonic/gin"
)

func DashboardRoute(router *gin.Engine, h *controllers.Handler) {
	router.GET("/dashboard", h.Dashboard)
}

