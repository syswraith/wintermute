package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)



func main() {

	r := gin.Default()

	r.GET("/:shortURL", func(c *gin.Context) {

		if c.Params.ByName("shortURL") == "whoami" {
			c.Redirect(http.StatusTemporaryRedirect, "https://syswraith.com")
			return
		}

		c.Redirect(http.StatusTemporaryRedirect, "https://youtu.be/WVZ6r-ld52k")
		
	})

	r.Run(":31337")

}
