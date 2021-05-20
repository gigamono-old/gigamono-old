package auth_test

import (
	"testing"

	"github.com/gigamono/gigamono/pkg/auth"
	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/stretchr/testify/assert"
)

func TestPasswordHashMatch(t *testing.T) {
	password := "password1234"

	t.Log(">> Hashing")
	hash, err := auth.GeneratePasswordHash(password, 10)
	assert.Nil(t, err)

	t.Log(">> Matching")
	err = auth.VerifyPasswordHash("password1234", hash)

	assert.Nil(t, err)
}

func TestPasswordHashMismatch(t *testing.T) {
	password := "password1234"

	t.Log(">> Hashing")
	hash, err := auth.GeneratePasswordHash(password, 10)
	assert.Nil(t, err)

	t.Log(">> Matching")
	err = auth.VerifyPasswordHash("password123", hash)

	assert.Equal(t, errs.NewTamperError(), err)
}
