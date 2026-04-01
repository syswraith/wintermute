package routes

import (
	"vulcan/internal/controllers"

	"github.com/gin-gonic/gin"
)

func ShortURLRoute(router *gin.Engine, h *controllers.Handler) {
	router.GET("/:shortURL", h.ShortenURL)
}

