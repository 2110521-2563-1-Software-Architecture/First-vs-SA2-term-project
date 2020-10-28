package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Key struct {
	key string
}

func GetUnusedKey(c *gin.Context) {
	// TODO: Gen keys from database.
	key := Key{ key: "1234" } 
	err := c.BindJSON(&key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(key)
	c.JSON(http.StatusOK, gin.H{ "key": key.key })
}
