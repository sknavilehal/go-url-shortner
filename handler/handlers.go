package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sknavilehal/go-url-shortener/shortner"
	"github.com/sknavilehal/go-url-shortener/store"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {

	var json UrlCreationRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	shortUrl := shortner.GenerateShortLink(json.LongUrl, json.UserId)

	store.SaveUrlMapping(shortUrl, json.LongUrl, json.UserId)

	host := "http://localhost:8080/"
	c.JSON(200, gin.H{
		"message":   "Short url created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")

	initalUrl := store.RetrieveInitialUrl(shortUrl)

	c.Redirect(302, initalUrl)
}
