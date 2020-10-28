package keygen

import (
	// "fmt"
	// "time"
	"github.com/2110521-2563-1-Software-Architecture/First-vs-SA2-term-project/repositories"
	"strconv"
)

func GenerateKeys(repo repositories.KeyRepository) {
	// TODO: Generate unique keys and insert into the database.
	for i := 0; i < 100; i++ {
		repo.InsertKey(strconv.Itoa(i))
	}	
}