package auth

import (
	"crypto/hmac"
	"crypto/sha512"

	"github.com/gigamono/gigamono/pkg/encodings"
	"github.com/gigamono/gigamono/pkg/errs"
)

// GenerateSignedCSRFID generates a signed and hashed CSRF ID from its plaintext form using provided secret key.
//
// This uses HMAC with SHA512 for hashing.
func GenerateSignedCSRFID(plaintextCSRFID string, secretKey []byte) (string, error) {
	mac := hmac.New(sha512.New, secretKey)
	_, err := mac.Write([]byte(plaintextCSRFID))
	if err != nil {
		return "", err
	}

	hash := mac.Sum(nil)

	return encodings.Base64URLEncode(hash), nil
}

// VerifySignedCSRFID verfies that hashed/signed CSRF ID was generated from the plaintext CSRF ID using specified secret key.
func VerifySignedCSRFID(plaintextCSRFID string, hashedCSRFID string, secretKey []byte) error {
	hash, err := GenerateSignedCSRFID(plaintextCSRFID, secretKey)
	if err != nil {
		return err
	}

	if hash != hashedCSRFID {
		return errs.NewTamperError()
	}

	return nil
}
