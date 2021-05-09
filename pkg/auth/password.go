package auth

import (
	"github.com/alexedwards/argon2id"
)

// GeneratePasswordHash generates a key from a password.
//
// This uses Argon2id key derivation function.
//
// https://en.wikipedia.org/wiki/Argon2
func GeneratePasswordHash(password string, iterations uint32) (string, error) {
	params := &argon2id.Params{
		Memory:      64 * 1024,
		Iterations:  iterations,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	return argon2id.CreateHash(password, params)
}

// VerifyPasswordHash verfies that hashed password is derived from plain text password.
func VerifyPasswordHash(plaintextPassword string, passwordHash string) (bool, error) {
	return argon2id.ComparePasswordAndHash(plaintextPassword, passwordHash)
}
