package routes

import (
	"vulcan/internal/controllers"

	"github.com/gin-gonic/gin"
)

func CreateRoute(router *gin.Engine, c *controllers.Control) {
	router.POST("/create", c.CreateLink)
}

