package routes

import (
	"vulcan/internal/middleware"

	"github.com/gin-gonic/gin"
)

func LoginRoute(router *gin.Engine) {
	router.POST("/login", middleware.LoginHandler)
}
