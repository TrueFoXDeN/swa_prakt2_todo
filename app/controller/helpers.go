package controller

import (
	"crypto/rand"
	"encoding/base64"
)

// Helper function
// Generates random Base64-URL encoded string composed of n bytes
func generateRandomString(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)

	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), err
}
