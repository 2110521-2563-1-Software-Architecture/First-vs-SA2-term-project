package main

import (
	"github.com/2110521-2563-1-Software-Architecture/First-vs-SA2-term-project/handlers"
	"github.com/2110521-2563-1-Software-Architecture/First-vs-SA2-term-project/keygen"
	"github.com/2110521-2563-1-Software-Architecture/First-vs-SA2-term-project/repositories"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := repositories.NewMemoryKeyRepository()

	go func() {
		keygen.GenerateKeys(repo)
	}()

	router := gin.Default()
	router.Use(cors.Default())
	router.Use(Repo(repo))
	router.GET("/", handlers.GetUnusedKey)
	router.Run(":8081")
}

func Repo(repo repositories.KeyRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
			c.Set("repo", repo)
			c.Next()
	}
}

