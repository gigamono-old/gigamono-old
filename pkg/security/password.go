package security

import (
	"github.com/alexedwards/argon2id"
	"github.com/gigamono/gigamono/pkg/errs"
)

// GeneratePasswordHash generates a password hash from its plaintext form.
//
// This uses Argon2id key derivation function.
//
// https://en.wikipedia.org/wiki/Argon2
func GeneratePasswordHash(plainTextPassword string, iterations uint32) (string, error) {
	// Sec: Beware of DoS.
	params := &argon2id.Params{
		Memory:      64 * 1024,
		Iterations:  iterations,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	return argon2id.CreateHash(plainTextPassword, params)
}

// VerifyPasswordHash verfies that password hash was generated from the plaintext password.
func VerifyPasswordHash(plaintextPassword string, passwordHash string) error {
	match, err := argon2id.ComparePasswordAndHash(plaintextPassword, passwordHash)
	if err != nil {
		return err
	}

	if !match {
		return errs.NewTamperError()
	}

	return nil
}
