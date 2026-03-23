package routes

import (
	"vulcan/internal/controllers"

	"github.com/gin-gonic/gin"
)

func DashboardRoute(router *gin.Engine, c *controllers.Control) {
	router.GET("/dashboard", c.Dashboard)
}

