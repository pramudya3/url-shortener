package helper

import (
	"math/rand"
	"time"
)

func RandString(n int) string {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"
	var randseed *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, n)
	for i := range b {
		b[i] = charset[randseed.Intn(len(charset))]
	}

	return string(b)
}
