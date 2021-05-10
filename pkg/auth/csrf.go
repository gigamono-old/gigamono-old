package auth

import (
	"crypto/hmac"
	"crypto/sha512"

	"github.com/gigamono/gigamono/pkg/encodings"
)

// GenerateSignedCSRFToken generates a signed CSRF token.
//
// This uses HMAC with SHA512 for hashing.
func GenerateSignedCSRFToken(csrfToken string, secretKey []byte) (string, error) {
	mac := hmac.New(sha512.New, secretKey)
	_, err := mac.Write([]byte(csrfToken))
	if err != nil {
		return "", err
	}

	hash := mac.Sum(nil)

	return encodings.NewBase64String(hash), nil
}

// VerifySignedCSRFToken verfies that signed CSRF token was generated using specified secret key.
func VerifySignedCSRFToken(csrfToken string, hashedCsrfToken string, secretKey []byte) (bool, error) {
	hash, err := GenerateSignedCSRFToken(csrfToken, secretKey)
	if err != nil {
		return false, err
	}

	return hash == hashedCsrfToken, nil
}
