package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/2110521-2563-1-Software-Architecture/First-vs-SA2-term-project/repositories"
)

type Key struct {
	key string
}

func GetUnusedKey(c *gin.Context) {
	// TODO: Gen keys from database.
	repo := c.MustGet("repo").(repositories.KeyRepository)
	newKey, err := repo.GetUnusedKey()
	key := Key{ key: newKey } 
	err = c.BindJSON(&key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(key)
	c.JSON(http.StatusOK, gin.H{ "key": key.key })
}
