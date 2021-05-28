package security

import (
	"testing"

	"github.com/gigamono/gigamono/pkg/security"
	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/stretchr/testify/assert"
)

func TestPasswordHashMatch(t *testing.T) {
	password := "password1234"

	t.Log(">> Hashing")
	hash, err := security.GeneratePasswordHash(password, 10)
	assert.Nil(t, err)

	t.Log(">> Matching")
	err = security.VerifyPasswordHash("password1234", hash)

	assert.Nil(t, err)
}

func TestPasswordHashMismatch(t *testing.T) {
	password := "password1234"

	t.Log(">> Hashing")
	hash, err := security.GeneratePasswordHash(password, 10)
	assert.Nil(t, err)

	t.Log(">> Matching")
	err = security.VerifyPasswordHash("password123", hash)

	assert.Equal(t, errs.NewTamperError(), err)
}
