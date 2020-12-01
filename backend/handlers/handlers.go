package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"regexp"

	"github.com/2110521-2563-1-Software-Architecture/First-vs-SA2-term-project/repositories"
	"github.com/2110521-2563-1-Software-Architecture/First-vs-SA2-term-project/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/vicanso/go-axios"
)

type ShortenURLPayload struct {
	URL string
}

type VisitRecord struct {
	Ip        string
	Hash      string
	Timestamp string
}

type Key struct {
	Key string
}

func GetUnusedKey() (string, error) {
	keygen_endpoint := fmt.Sprintf("http://%s:%s", utils.Getenv("KEYGEN_HOST"), utils.Getenv("KEYGEN_PORT"))
	resp, err := axios.Get(keygen_endpoint)
	if err != nil {
		return "", errors.New("Error while retrieving new key")
	}
	var k Key
	json.Unmarshal(resp.Data, &k)
	return k.Key, nil
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

	re := regexp.MustCompile(`^(?:f|ht)tps?\:\/\/.*`)
	url := body.URL
	if re.FindString(url) == "" {
		url = "http://" + url
	}
	fmt.Println(url)

	_, err = repo.Create(hash, url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"url": url, "key": hash})
	}
}

func Redirect(c *gin.Context) {
	redis := c.MustGet("redis").(*redis.Client)
	repo := c.MustGet("repo").(repositories.URLRepository)
	repoHistory := c.MustGet("repoHistory").(repositories.HistoryRepository)

	hash := c.Param("hash")
	fmt.Println("hash: " + hash)

	// Insert record into the database
	exists, _ := repo.Exists(hash)
	if exists {
		visitRecord := VisitRecord{Hash: "hash-" + hash, Ip: "ip-" + c.ClientIP(), Timestamp: "time-" + time.Now().String()}
		fmt.Println(visitRecord)
		repoHistory.CreateHistory(c.ClientIP(), hash, time.Now().String())
	}

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
	fmt.Println("location: " + location)
}

func ShortenHistory(c *gin.Context) {
	repoHistory := c.MustGet("repoHistory").(repositories.HistoryRepository)

	var body Key
	err := c.BindJSON(&body)
	if err != nil {
		fmt.Println(err)
	}

	history, err := repoHistory.GetHistory(body.Key)
	fmt.Println("ShortenHistory:", history)
	c.JSON(http.StatusOK, string(history))
}
