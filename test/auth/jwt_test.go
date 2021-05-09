package auth_test

import (
	"testing"

	"github.com/gigamono/gigamono/pkg/auth"
	"github.com/stretchr/testify/assert"
)

const privKey = `
-----BEGIN EC PRIVATE KEY-----
MIHcAgEBBEIBVhNtM/d/7Xze73+9MpqJ5OZpLga11HdViL58j5mY7AZULh2EZo3V
BbLhqQzCdwNzg7PkyZDRPx65B9ne3ZCAOzagBwYFK4EEACOhgYkDgYYABAHOgyqG
pGs3wAbBzO5yLwB4q+J0lmhB93CEVymPX6IRxYoE5AMzZea+cuyVJA27UcTxwCIR
DAF7sC+b72s4QIn5HQH4ZAIrcJy1Daw/ZdZpcUF9zFAJG65Bpqh0nTnF8M4fDW1W
5g5N14c1p+fxsMhiNSLD/yEp0vtoqA8uiRS0eIg4eg==
-----END EC PRIVATE KEY-----
`

const validPubKey = `
-----BEGIN PUBLIC KEY-----
MIGbMBAGByqGSM49AgEGBSuBBAAjA4GGAAQBzoMqhqRrN8AGwczuci8AeKvidJZo
QfdwhFcpj1+iEcWKBOQDM2XmvnLslSQNu1HE8cAiEQwBe7Avm+9rOECJ+R0B+GQC
K3CctQ2sP2XWaXFBfcxQCRuuQaaodJ05xfDOHw1tVuYOTdeHNafn8bDIYjUiw/8h
KdL7aKgPLokUtHiIOHo=
-----END PUBLIC KEY-----
`

const invalidPubKey = `
-----BEGIN PUBLIC KEY-----
MIGbMBAGByqGSM49AgEGBSuBBAAjA4GGAAQB0lv69hDz4SjRFX4R4V3zwXUrCBqc
vqYP5IWj6f1hj95IFDqEosJnYcDpT7Ot1op21lmRRcog4EcumMbmTMj7Ay4AOUuW
REXMDBIwoLeGnuJMRi2BbCNer5EY+/gkgyLqjE2CHcTLO9Sn2BOjOvjZMZQ6WyS0
YNJ5qPakZwr7GlDprxc=
-----END PUBLIC KEY-----
`

func TestValidPublicKey(t *testing.T) {
	payload := auth.GenerateAuthClaims("subject_id", "signup", 604800)

	t.Log(">> Signing")
	token, err := auth.GenerateJWTToken(payload, []byte(privKey))
	assert.Nil(t, err)

	t.Log(">> token =", token)

	t.Log(">> Decoding")
	claims, err := auth.DecodeJWTToken(token, []byte(validPubKey))
	assert.Nil(t, err)

	assert.Equal(t, &payload, claims)
}

func TestInvalidPublicKey(t *testing.T) {
	payload := auth.GenerateAuthClaims("subject_id", "signup", 604800)

	t.Log(">> Signing")
	token, err := auth.GenerateJWTToken(payload, []byte(privKey))
	assert.Nil(t, err)

	t.Log(">> token =", token)

	t.Log(">> Decoding")
	claims, err := auth.DecodeJWTToken(token, []byte(invalidPubKey))
	assert.NotNil(t, err)

	t.Log(">> err", err)

	assert.Equal(t, &auth.Claims{}, claims)
}
