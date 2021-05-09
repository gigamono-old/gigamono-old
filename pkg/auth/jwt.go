package auth

import (
	"encoding/json"
	"time"

	jose "github.com/dvsekhvalnov/jose2go"
	"github.com/dvsekhvalnov/jose2go/keys/ecc"
)

// StandardClaims standard claims
type StandardClaims struct {
	Issuer    string `json:"iss,omitempty"`
	Subject   string `json:"sub,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
}

// Claims is a custom claims type wrapping JWT standard claims.
type Claims struct {
	StandardClaims
	Action string `json:"action,omitempty"`
}

// GenerateSignedJWT generates a JWT token from payload signed by a private key.
//
// This uses an ECDSA P-521 asymmetric encryption with SHA-512 hashing.
//
// https://en.wikipedia.org/wiki/Elliptic-curve_cryptography
func GenerateSignedJWT(payload Claims, privateKeyBytes []byte) (string, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	payloadString := string(payloadBytes)
	privateKey, err := ecc.ReadPrivate(privateKeyBytes)
	if err != nil {
		return "", err
	}

	return jose.Sign(payloadString, jose.ES512, privateKey, jose.Header("typ", "JWT"))
}

// GenerateAuthClaims generates JWT claims for auth server.
func GenerateAuthClaims(subject string, action string, expirationInSeconds int) Claims {
	now := time.Now()

	return Claims{
		StandardClaims: StandardClaims{
			Issuer:    "auth.gigamono.com",
			Subject:   subject,
			ExpiresAt: now.Add(time.Second * time.Duration(expirationInSeconds)).Unix(),
			IssuedAt:  now.Unix(),
		},
		Action: action,
	}
}

// DecodeSignedJWT verifies that the token was signed with associated private key.
func DecodeSignedJWT(tokenString string, publicKeyBytes []byte) (*Claims, error) {
	publicKey, err := ecc.ReadPublic(publicKeyBytes)
	if err != nil {
		return &Claims{}, err
	}

	payload, _, err := jose.Decode(tokenString, publicKey)
	if err != nil {
		return &Claims{}, err
	}

	var claims Claims
	err = json.Unmarshal([]byte(payload), &claims)

	return &claims, err
}
