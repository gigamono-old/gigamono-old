package auth

import "crypto/rand"

// GenerateRandomBytes generates a cryptograhically secure random vector of bytes with specified length.
func GenerateRandomBytes(length uint) ([]byte, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}
