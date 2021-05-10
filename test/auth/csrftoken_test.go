package auth

import (
	"testing"

	"github.com/gigamono/gigamono/pkg/auth"
	"github.com/stretchr/testify/assert"
)

var secretKey = []byte("secret-key-of-the-secret-mission")

func TestCSRFTokenHashMatch(t *testing.T) {
	csrfToken := "fgHhhonnZ_FNBrZQfCa99A"

	t.Log(">> Hashing")
	hash, err := auth.GenerateSignedCSRFToken(csrfToken, secretKey)
	assert.Nil(t, err)

	t.Log(">> Matching")
	ok, err := auth.VerifySignedCSRFToken("fgHhhonnZ_FNBrZQfCa99A", hash, secretKey)
	assert.Nil(t, err)

	assert.Equal(t, true, ok)
}

func TestCSRFTokenHashMismatch(t *testing.T) {
	csrfToken := "fgHhhonnZ_FNBrZQfCa99A"

	t.Log(">> Hashing")
	hash, err := auth.GenerateSignedCSRFToken(csrfToken, secretKey)
	assert.Nil(t, err)

	t.Log(">> Matching")
	ok, err := auth.VerifySignedCSRFToken("fgHhhonnZ_FNBrZQfCa99B", hash, secretKey)
	assert.Nil(t, err)

	assert.Equal(t, false, ok)
}
