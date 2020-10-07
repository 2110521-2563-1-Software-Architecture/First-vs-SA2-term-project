package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	URL string
}

func shortenURL(c *gin.Context) {
	// TODO add logic for create random string and store in databse
	// var body requestBody
	// err := c.ShouldBindBodyWith(binding.JSON)
	var body RequestBody
	err := c.BindJSON(&body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(body)
	c.JSON(http.StatusOK, gin.H{"url": body.URL})
}
func redirect(c *gin.Context) {
	hash := c.Param("hash")
	fmt.Println(hash)
	// TODO read from database and redirect
	const location = "https://google.co.th"
	c.Redirect(301, location)
}
func main() {
	fmt.Println("Hello world")
	router := gin.Default()
	router.POST("/shorten", shortenURL)
	router.GET("/:hash", redirect)
	router.Run(":8080")
}
