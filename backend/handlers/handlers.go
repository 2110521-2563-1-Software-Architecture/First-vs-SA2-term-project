package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"errors"
	"json"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/vicanso/go-axios"
	"github.com/2110521-2563-1-Software-Architecture/First-vs-SA2-term-project/repositories"
	"github.com/2110521-2563-1-Software-Architecture/First-vs-SA2-term-project/utils"
)

type ShortenURLPayload struct {
	URL string
}

type VisitRecord struct {
	Ip        string
	Hash      string
	Timestamp string
}

type URLMap struct {
	Key string
	Url string
}

func GetUnusedKey() (string, error) {
	keygen_endpoint := fmt.Sprintf("http://%s:%s", utils.Getenv("KEYGEN_HOST"), utils.Getenv("KEYGEN_PORT"))
	resp, err := axios.Get(keygen_endpoint)
	if err != nil {
		return "", errors.New("Error while retrieving new key")
	}
	// var url_map URLMap
	// url_map = json
	return string(resp.Data), nil
}

func ShortenURL(c *gin.Context) {
	repo := c.MustGet("repo").(repositories.URLRepository)

	var body ShortenURLPayload
	err := c.BindJSON(&body)
	if err != nil {
		fmt.Println(err)
	}

	hash, err := GetUnusedKey()
	if err != nil {
		fmt.Println(err)
	}

	_, err = repo.Create(hash, body.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"url": body.URL, "key": hash})
	}
}

func Redirect(c *gin.Context) {
	redis := c.MustGet("redis").(*redis.Client)
	repo := c.MustGet("repo").(repositories.URLRepository)

	hash := c.Param("hash")

	// TODO insert ip record into the database
	visitRecord := VisitRecord{Hash: hash, Ip: c.ClientIP(), Timestamp: time.Now().String()}
	fmt.Println(visitRecord)

	// Read from cache first
	location, err := redis.Get(redis.Context(), hash).Result()
	if err != nil {
		fmt.Println("Cache not found")
		// Cache not found, query in db instead.
		location, err = repo.GetURL(hash)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			redis.Set(redis.Context(), hash, location, 0).Result()
			c.Redirect(301, location)
		}
	} else {
		fmt.Println("Cache found")
		c.Redirect(301, location)
	}
	fmt.Println(location)
}

func ShortenHistory(c *gin.Context) {
	//TODO assign value from DB & cast go struct to JSON!!!!
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
