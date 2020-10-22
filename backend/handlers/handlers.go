package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShortenURLPayload struct {
	URL string
}

func ShortenURL(c *gin.Context) {
	var body ShortenURLPayload
	err := c.BindJSON(&body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(body)
	c.JSON(http.StatusOK, gin.H{"url": body.URL})
}

func Redirect(c *gin.Context) {
	hash := c.Param("hash")
	fmt.Println(hash)
	// TODO read from database and redirect
	location := "https://google.co.th"
	c.Redirect(301, location)
}
