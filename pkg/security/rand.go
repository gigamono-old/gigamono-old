package security

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateRandomBase64 generates a cryptograhically secure random vector of bytes
// of specified encoded as base64.
func GenerateRandomBase64(length uint) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}
