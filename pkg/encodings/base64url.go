package encodings

import (
	"encoding/base64"
	"strings"
)

// Base64URLEncode creates a URL-encoded base64 string from a vector of bytes.
func Base64URLEncode(bytes []byte) string {
	// Create Base64 encoding from bytes.
	b64String := base64.RawStdEncoding.Strict().EncodeToString(bytes)

	// URL-encode Base64 string.
	b64String = strings.Replace(b64String, "+", "-", -1) // 62nd char of encoding
	b64String = strings.Replace(b64String, "/", "_", -1) // 63rd char of encoding

	return b64String
}

// Base64URLDecode decodes a URL-encoded base64 string.
func Base64URLDecode(b64String string) ([]byte, error) {
	// Replace URL-encodings.
	b64String = strings.Replace(b64String, "-", "+", -1) // 62nd char of encoding
	b64String = strings.Replace(b64String, "_", "/", -1) // 63rd char of encoding

	return base64.RawStdEncoding.Strict().DecodeString(b64String)
}

