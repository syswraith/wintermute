package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func main() {

	r := gin.Default()


	r.GET("/", func(x *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "https://ctrl-c.club/~fey/")
	})


	r.GET("/:shortURL", func(c *gin.Context) {

		if c.Params.ByName("shortURL") == "whoami" {
			c.Redirect(http.StatusTemporaryRedirect, "https://syswraith.com")
			return
		}

		c.Redirect(http.StatusTemporaryRedirect, "https://youtu.be/O9BK3xcRH1g")
		
	})

	r.Run(":31337")

}
