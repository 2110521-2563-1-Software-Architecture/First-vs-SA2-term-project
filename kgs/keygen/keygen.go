package keygen

import (
	"fmt"
	"time"
)

func GenerateKeys() {
	// TODO: Generate unique keys and insert into the database.
	for i := 0; i < 100; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("Key")
	}
}