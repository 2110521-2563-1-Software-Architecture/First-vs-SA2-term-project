package main

import (
	"github.com/2110521-2563-1-Software-Architecture/First-vs-SA2-term-project/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/shorten", handlers.ShortenURL)
	router.GET("/:hash", handlers.Redirect)
	router.Run(":8080")
}
