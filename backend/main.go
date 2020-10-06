package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectMongo() (c *Client) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://url-matcher:EeyFLrO0PLRrqXDe@cluster0.kj1oc.gcp.mongodb.net/url-matcher?retryWrites=true&w=majority"))
	if err != nil {
		fmt.Println(err)
		// return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	return client
	client.Database('url-matcher').
}

type RequestBody struct {
	URL string `json:"url"`
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
