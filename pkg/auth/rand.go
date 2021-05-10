package auth

import (
	"crypto/rand"

	"github.com/gigamono/gigamono/pkg/encodings"
)

// GenerateRandomBase64 generates a cryptograhically secure random vector of bytes
// of specified encoded as base64.
func GenerateRandomBase64(length uint) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return encodings.NewBase64String(bytes), nil
}
