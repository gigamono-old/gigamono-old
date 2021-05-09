package auth

import (
	"testing"

	"github.com/gigamono/gigamono/pkg/auth"
	"github.com/stretchr/testify/assert"
)

var secretKey = []byte("secret-key-of-the-secret-mission")

func TestCSRFTokenHashMatch(t *testing.T) {
	csrfToken, err := auth.GenerateRandomBytes(16)
	csrfTokenString := string(csrfToken)
	assert.Nil(t, err)

	t.Log(">> Hashing")
	hash, err := auth.GenerateSignedCSRFToken(csrfTokenString, secretKey)
	assert.Nil(t, err)

	t.Log(">> Matching")
	ok, err := auth.VerifySignedCSRFToken(csrfTokenString, hash, secretKey)
	assert.Nil(t, err)

	assert.Equal(t, true, ok)
}

func TestCSRFTokenHashMismatch(t *testing.T) {
	csrfToken, err := auth.GenerateRandomBytes(16)
	csrfTokenString := string(csrfToken)
	assert.Nil(t, err)

	t.Log(">> Hashing")
	hash, err := auth.GenerateSignedCSRFToken(csrfTokenString, secretKey)
	assert.Nil(t, err)

	t.Log(">> Matching")
	ok, err := auth.VerifySignedCSRFToken("123456789ABCDEF", hash, secretKey)
	assert.Nil(t, err)

	assert.Equal(t, false, ok)
}
