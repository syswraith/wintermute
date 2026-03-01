package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"vulcan/minerva"
)


func main() {

	r := gin.Default()


	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "https://ctrl-c.club/~fey/")
	})

	r.POST("/create", func(c *gin.Context) {

	})


	r.GET("/:shortURL", func(c *gin.Context) {

		if c.Params.ByName("shortURL") == "whoami" {
			c.Redirect(http.StatusTemporaryRedirect, "https://syswraith.com")
			return
		}

		c.Redirect(http.StatusTemporaryRedirect, "https://youtu.be/O9BK3xcRH1g")

	})

	r.Run(":31337")


	// testing db

	db := minerva.Connect()
	minerva.Ping(db)
	minerva.SelectAll(db)
	db.Close()
}
