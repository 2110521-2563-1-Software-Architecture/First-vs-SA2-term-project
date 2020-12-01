package main

import (
	"github.com/2110521-2563-1-Software-Architecture/First-vs-SA2-term-project/handlers"
	"github.com/2110521-2563-1-Software-Architecture/First-vs-SA2-term-project/repositories"
	"github.com/2110521-2563-1-Software-Architecture/First-vs-SA2-term-project/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: utils.Getenv("REDIS_HOST") + ":" + utils.Getenv("REDIS_PORT"),
	})

	return client
}

func main() {
	rClient := RedisClient()
	repo := repositories.NewMongoURLRepository()
	repoHistory := repositories.NewMongoHistoryRepository()

	router := gin.Default()

	router.Use(cors.Default())
	router.Use(Redis(rClient))
	router.Use(Repo(repo))
	router.Use(RepoHistory(repoHistory))

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

func Repo(repo repositories.URLRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("repo", repo)
		c.Next()
	}
}

func RepoHistory(repoHistory repositories.HistoryRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("repoHistory", repoHistory)
		c.Next()
	}
}
