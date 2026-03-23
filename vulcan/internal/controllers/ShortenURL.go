package controllers

import (
	"net/http"
	"vulcan/internal/repository"

	"github.com/gin-gonic/gin"
)

func (c *Control) ShortenURL(context *gin.Context) {
	shortURL := context.Params.ByName("shortURL")

	if shortURL == "whoami" {
		context.Redirect(http.StatusTemporaryRedirect, "https://syswraith.com")
		return
	}

	longURL, err := repository.Fetch(shortURL, c.DB)
	if err != nil {
		if err.Error() == "link expired" {
			context.JSON(http.StatusGone, gin.H{"error": "link expired"})
			return
		}

		context.JSON(http.StatusNotFound, gin.H{"error": "page not found"})
		return
	}
	context.Redirect(http.StatusTemporaryRedirect, longURL)
}

