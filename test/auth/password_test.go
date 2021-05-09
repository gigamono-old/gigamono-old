package auth_test

import (
	"testing"

	"github.com/gigamono/gigamono/pkg/auth"
	"github.com/stretchr/testify/assert"
)

func TestHashMatch(t *testing.T) {
	password := "password1234"

	t.Log(">> Hashing")
	hash, err := auth.GenerateHash(password, 10)
	assert.Nil(t, err)

	t.Log(">> Matching")
	ok, err := auth.VerifyHash("password1234", hash)
	assert.Nil(t, err)

	assert.Equal(t, true, ok)
}

func TestHashMismatch(t *testing.T) {
	password := "password1234"

	t.Log(">> Hashing")
	hash, err := auth.GenerateHash(password, 10)
	assert.Nil(t, err)

	t.Log(">> Matching")
	ok, err := auth.VerifyHash("password123", hash)
	assert.Nil(t, err)

	assert.Equal(t, false, ok)
}
