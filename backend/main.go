package main

import (
	"github.com/2110521-2563-1-Software-Architecture/First-vs-SA2-term-project/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/2110521-2563-1-Software-Architecture/First-vs-SA2-term-project/utils"
)

func RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: utils.Getenv("REDIS_HOST") + ":" + utils.Getenv("REDIS_PORT"),
	})
	
	return client
}

func main() {
	rClient := RedisClient()

	router := gin.Default()
	router.Use(cors.Default())
	router.Use(Redis(rClient))
	router.POST("/shorten", handlers.ShortenURL)
	router.POST("/shortenHistory", handlers.ShortenHistory)
	router.GET("/:hash", handlers.Redirect)
	router.Run(":8080")
}

func Redis(client *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
			c.Set("redis", client)
			c.Next()
	}
}

