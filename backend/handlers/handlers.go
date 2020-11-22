package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
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
	redis := c.MustGet("redis").(*redis.Client)
	hash := c.Param("hash")

	// TODO insert ip record into the database
	visitRecord := VisitRecord{ hash: hash, ip: c.ClientIP(), timestamp: time.Now().String() }
	fmt.Println(visitRecord)
	
	// Read from cache first
	location, err := redis.Get(redis.Context(), hash).Result()
	if err != nil {
		fmt.Println("Cache not found")
		// Cache not found, query in db instead.
		location = "https://google.co.th"
		redis.Set(redis.Context(), hash, location, 0).Result()
	} else {
		fmt.Println("Cache found")
	}

	c.Redirect(301, location)
}
