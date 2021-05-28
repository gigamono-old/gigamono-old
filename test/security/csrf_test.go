package security

import (
	"testing"

	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/gigamono/gigamono/pkg/security"
	"github.com/stretchr/testify/assert"
)

var secretKey = []byte("secret-key-of-the-secret-mission")

func TestCSRFTokenHashMatch(t *testing.T) {
	csrfID := "fgHhhonnZ_FNBrZQfCa99A"

	t.Log(">> Hashing")
	hash, err := security.GenerateSignedCSRFID(csrfID, secretKey)
	assert.Nil(t, err)

	t.Log(">> Matching")
	err = security.VerifySignedCSRFID("fgHhhonnZ_FNBrZQfCa99A", hash, secretKey)

	assert.Nil(t, err)
}

func TestCSRFTokenHashMismatch(t *testing.T) {
	csrfID := "fgHhhonnZ_FNBrZQfCa99A"

	t.Log(">> Hashing")
	hash, err := security.GenerateSignedCSRFID(csrfID, secretKey)
	assert.Nil(t, err)

	t.Log(">> Matching")
	err = security.VerifySignedCSRFID("fgHhhonnZ_FNBrZQfCa99B", hash, secretKey)

	assert.Equal(t, errs.NewTamperError(), err)
}
