package keygen

import (
	"github.com/2110521-2563-1-Software-Architecture/First-vs-SA2-term-project/repositories"
	"math/rand"
	"time"
)
  
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
  
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))
  
func RandomStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
  
func RandomString(length int) string {
	return RandomStringWithCharset(length, charset)
}

func GenerateKeys(repo repositories.KeyRepository) {
	// TODO: Generate unique keys and insert into the database.
	for i := 0; i < 1000; i++ {
		repo.InsertKey(RandomString(4))
	}	
}