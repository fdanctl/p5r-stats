package utils

import (
	"crypto/rand"
	"encoding/hex"
)

// RandomID generates a random ID of length n using cryptographically secure randomness.
func RandomID(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
