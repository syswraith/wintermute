package routes

import (
	"vulcan/internal/controllers"

	"github.com/gin-gonic/gin"
)

func ShortURLRoute(router *gin.Engine, c *controllers.Control) {
	router.GET("/:shortURL", c.ShortenURL)
}

