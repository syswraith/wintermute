package main

import (
	"log"
	"net/http"
	"vulcan/minerva"

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
				"error":err,
			})
		}

		shortURL, err := minerva.Create(json.LongURL, db)

		if err != nil {
			log.Fatal("link generation failed")

		}

		shortURL = "https://localhost:31337/" + shortURL

		context.JSON(200, gin.H{"data": shortURL,})

	})


	router.GET("/:shortURL", func(context *gin.Context) {

		if context.Params.ByName("shortURL") == "whoami" {
			context.Redirect(http.StatusTemporaryRedirect, "https://syswraith.com")
			return
		}

		context.Redirect(http.StatusTemporaryRedirect, "https://youtu.be/O9BK3xcRH1g")

	})

	// run on 31337 because we elite B)
	router.Run(":31337")

}
