package auth

import (
	"encoding/json"
	"time"

	jose "github.com/dvsekhvalnov/jose2go"
	"github.com/dvsekhvalnov/jose2go/keys/ecc"
	"github.com/gigamono/gigamono/pkg/errs"
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
	SignedCSRFID string `json:"signed_csrf_id,omitempty"`
	Action       string `json:"action,omitempty"`
}

// GenerateSignedJWT generates a JWT token from payload signed by a private key.
//
// This uses an ECDSA P-521 asymmetric encryption with SHA-512 hashing.
//
// https://en.wikipedia.org/wiki/Elliptic-curve_cryptography
func GenerateSignedJWT(payload *Claims, privateKeyBytes []byte) (string, error) {
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

// GeneratePreSessionClaims generates JWT claims for a pre-session user.
func GeneratePreSessionClaims(signedCSRFID string, expirationInSeconds int) *Claims {
	now := time.Now()

	return &Claims{
		StandardClaims: StandardClaims{
			Issuer:    "auth.gigamono.com",
			Subject:   "pre-session",
			ExpiresAt: now.Add(time.Second * time.Duration(expirationInSeconds)).Unix(),
			IssuedAt:  now.Unix(),
		},
		SignedCSRFID: signedCSRFID,
		Action:       "pre-session",
	}
}

// GenerateSessionClaims generates JWT claims for a session user.
func GenerateSessionClaims(subject string, signedCSRFID string, action string, expirationInSeconds int) *Claims {
	now := time.Now()

	return &Claims{
		StandardClaims: StandardClaims{
			Issuer:    "auth.gigamono.com",
			Subject:   subject,
			ExpiresAt: now.Add(time.Second * time.Duration(expirationInSeconds)).Unix(),
			IssuedAt:  now.Unix(),
		},
		SignedCSRFID: signedCSRFID,
		Action:       action,
	}
}

// DecodeAndVerifySignedJWT decodes and verifies that the token was signed with associated private key as well as still within expriration limit.
func DecodeAndVerifySignedJWT(tokenString string, publicKeyBytes []byte) (*Claims, error) {
	publicKey, err := ecc.ReadPublic(publicKeyBytes)
	if err != nil {
		return &Claims{}, err
	}

	// Decode payload.
	payload, _, err := jose.Decode(tokenString, publicKey)
	if err != nil {
		return &Claims{}, errs.NewTamperError()
	}

	// Deserialize claims.
	var claims Claims
	if err = json.Unmarshal([]byte(payload), &claims); err != nil {
		return &Claims{}, err
	}

	// Check that token has not expired.
	if claims.ExpiresAt < time.Now().Unix() {
		return &Claims{}, errs.NewExpirationError()
	}

	return &claims, nil
}
