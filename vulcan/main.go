package main

import (
	"log"
	"net/http"

	"vulcan/internal/controllers"
	"vulcan/internal/models"
	"vulcan/internal/repository"
	"vulcan/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// connect database
	repo := repository.Connect()
	db := repo.DB

	// initialize controllers
	c := controllers.Init(repo)

	// create da router
	router := gin.Default()

	err := db.AutoMigrate(&models.Link{})
	if err != nil {
		log.Fatal(err)
	}

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(config))

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "https://ctrl-c.club/~fey/")
	})

	// load routes
	routes.CreateRoute(router, c)
	routes.DashboardRoute(router, c)
	routes.ShortURLRoute(router, c)
	routes.LoginRoute(router)

	// run on 31337 because we elite B)
	router.Run(":31337")
}
