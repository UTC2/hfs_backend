package common

import (
	"crypto/rand"
	"math/big"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenSalt(length int) string {
	if length <= 0 {
		length = 50
	}
	b := make([]rune, length)
	max := big.NewInt(int64(len(letters)))
	for i := range b {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			panic(err) // crypto/rand.Reader should not fail; if it does, abort
		}
		b[i] = letters[n.Int64()]
	}
	return string(b)
}
