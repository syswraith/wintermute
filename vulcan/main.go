package main

import (
	"log"
	"net/http"
	"vulcan/minerva"

	"github.com/gin-gonic/gin"
)

type postJSON struct {
	Link string `json:"longURL" binding:"required"`
}



func main() {

	db := minerva.Connect()

	router := gin.Default()


	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "https://ctrl-c.club/~fey/")
	})

	router.POST("/create", func(context *gin.Context) {
		var json postJSON

		if err := context.ShouldBindJSON(&json); err != nil {
			context.JSON(400, gin.H{"error":"1"})
		}

		shortlink, err := minerva.Create(json.Link, db)

		if err != nil {
			log.Fatal("link generation failed")

		}

		shortlink = "https://localhost:31337/" + shortlink

		context.JSON(200, gin.H{"success":"1", "data": shortlink,})

	})


	router.GET("/:shortURL", func(context *gin.Context) {

		if context.Params.ByName("shortURL") == "whoami" {
			context.Redirect(http.StatusTemporaryRedirect, "https://syswraith.com")
			return
		}

		context.Redirect(http.StatusTemporaryRedirect, "https://youtu.be/O9BK3xcRH1g")

	})



	router.Run(":31337")

}
