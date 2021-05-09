package auth

import (
	"github.com/alexedwards/argon2id"
)

// GenerateHash uses Argon2id KDF to generate a key from password.
func GenerateHash(password string, iterations uint32) (string, error) {
	params := &argon2id.Params{
		Memory:      64 * 1024,
		Iterations:  iterations,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	return argon2id.CreateHash(password, params)
}

// VerifyHash verfies that hashed password is derived from plain text password.
func VerifyHash(plaintextPassword string, passwordHash string) (bool, error) {
	return argon2id.ComparePasswordAndHash(plaintextPassword, passwordHash)
}
