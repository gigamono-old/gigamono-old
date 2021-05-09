package auth

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"strings"
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

	// Create Base64 encoding from hash.
	b64hash := base64.RawStdEncoding.Strict().EncodeToString(hash)

	// URL-encode Base64 string.
	b64hash = strings.Replace(b64hash, "+", "-", -1) // 62nd char of encoding
	b64hash = strings.Replace(b64hash, "/", "_", -1) // 63rd char of encoding

	return b64hash, nil
}

// VerifySignedCSRFToken verfies that signed CSRF token was generated using specified secret key.
func VerifySignedCSRFToken(csrfToken string, hashedCsrfToken string, secretKey []byte) (bool, error) {
	hash, err := GenerateSignedCSRFToken(csrfToken, secretKey)
	if err != nil {
		return false, err
	}

	return hash == hashedCsrfToken, nil
}
