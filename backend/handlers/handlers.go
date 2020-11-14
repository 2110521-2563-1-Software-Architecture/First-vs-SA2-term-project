package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ShortenURLPayload struct {
	URL string
}

type VisitRecord struct {
	ip string
	hash string
	timestamp string
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
	// TODO insert ip record into the database
	visitRecord := VisitRecord{ hash: c.Param("hash"), ip: c.ClientIP(), timestamp: time.Now().String() }
	fmt.Println(visitRecord)
	// TODO read from database and redirect
	location := "https://google.co.th"
	c.Redirect(301, location)
}
