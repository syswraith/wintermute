package main

import (
	"log"
	"net/http"
	"vulcan/minerva"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type postJSON struct {
	LongURL string `json:"longURL" binding:"required"`
}

func main() {

	// connect database
	db := minerva.Connect()

	// create da router
	router := gin.Default()

	router.Use(cors.Default())

	// endpoints
	// `/:shortURL` redirect to shorturl
	// `/create` create shorturl from longurl

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "https://ctrl-c.club/~fey/")
	})

	router.POST("/create", func(context *gin.Context) {
		var json postJSON

		if err := context.ShouldBindJSON(&json); err != nil {
			context.JSON(400, gin.H{
				"error": err,
			})
		}

		shortURL, err := minerva.Create(json.LongURL, db)

		if err != nil {
			log.Fatal("link generation failed")

		}

		shortURL = "http://localhost:31337/" + shortURL

		context.JSON(http.StatusOK, gin.H{"shortURL": shortURL})

	})

	router.GET("/:shortURL", func(context *gin.Context) {
		shortURL := context.Params.ByName("shortURL")

		if shortURL == "whoami" {
			context.Redirect(http.StatusTemporaryRedirect, "https://syswraith.com")
			return
		}

		longURL, err := minerva.Fetch(shortURL, db)

		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": "page not found"})
		}

		context.Redirect(http.StatusTemporaryRedirect, longURL)

	})

	// run on 31337 because we elite B)
	router.Run(":31337")

}
