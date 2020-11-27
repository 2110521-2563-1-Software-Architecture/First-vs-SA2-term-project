package handlers

import (
	"encoding/json"
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
	Ip        string
	Hash      string
	Timestamp string
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
	visitRecord := VisitRecord{Hash: hash, Ip: c.ClientIP(), Timestamp: time.Now().String()}
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

func ShortenHistory(c *gin.Context) {
	//TODO assign value from DB & cast go struct to JSON!!!!

	// var history = []VisitRecord{
	// 	VisitRecord{
	// 		Ip:        "1.2.3",
	// 		Hash:      "goo.gl/1234",
	// 		Timestamp: "12354394584",
	// 	},
	// 	VisitRecord{
	// 		Ip:        "1.2.3",
	// 		Hash:      "goo.gl/1234",
	// 		Timestamp: "12354394584",
	// 	},
	// }

	var history [5]VisitRecord
	history[0] = VisitRecord{
		Ip:        "1.2.3",
		Hash:      "goo.gl/1234",
		Timestamp: "12354394584",
	}
	history[1] = VisitRecord{
		Ip:        "1.2.3",
		Hash:      "goo.gl/1234",
		Timestamp: "12354394584",
	}

	recordsMarshal, err := json.Marshal(history)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(recordsMarshal))

	c.JSON(http.StatusOK, string(recordsMarshal))
}
