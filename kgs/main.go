package main

import (
	"github.com/2110521-2563-1-Software-Architecture/First-vs-SA2-term-project/handlers"
	"github.com/2110521-2563-1-Software-Architecture/First-vs-SA2-term-project/keygen"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	go func() {
		keygen.GenerateKeys()
	}()

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/", handlers.GetUnusedKey)
	router.Run(":8081")
}
