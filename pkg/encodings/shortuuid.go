package encodings

import (
	"encoding/base64"

	"github.com/gofrs/uuid"
)

// ShortenUUID reduces UUID to a shorted non-padded base64 string.
func ShortenUUID(uuid uuid.UUID) string {
	str := uuid.String()
	return ShortenUUIDString(str)
}

// ShortenUUIDString reduces UUID to a shorted non-padded base64 string.
func ShortenUUIDString(uuidString string) string {
	bytes := []byte(uuidString)
	return ShortenUUIDBytes(bytes)
}

// ShortenUUIDBytes reduces UUID to a shorted non-padded base64 string.
func ShortenUUIDBytes(bytes []byte) string {
	// Return url-encoded non-padded base64 string.
	return base64.RawURLEncoding.EncodeToString(bytes)
}

// DecodeShortenedUUID decodes shortened UUID.
func DecodeShortenedUUID(b64String string) (uuid.UUID, error) {
	bytes, err := base64.RawURLEncoding.Strict().DecodeString(b64String)
	if err != nil {
		return uuid.UUID{}, err
	}

	return uuid.FromBytes(bytes)
}
