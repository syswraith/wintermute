package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ShortenURL(context *gin.Context) {
	shortURL := context.Params.ByName("shortURL")

	if shortURL == "whoami" {
		context.Redirect(http.StatusTemporaryRedirect, "https://syswraith.com")
		return
	}

	longURL, err := h.Repo.Fetch(shortURL)
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
