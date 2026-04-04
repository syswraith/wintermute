package routes

import (
	"vulcan/internal/controllers"
	"vulcan/internal/middleware"

	"github.com/gin-gonic/gin"
)

func DashboardRoute(router *gin.Engine, h *controllers.Handler) {
	router.GET("/dashboard", middleware.RequireAuth(), h.Dashboard)
}
