package encodings

import (
	"encoding/base64"
	"strings"
)

// NewBase64String creates a URL-encoded base64 string from a vector of bytes.
func NewBase64String(bytes []byte) string {
	// Create Base64 encoding from bytes.
	b64String := base64.RawStdEncoding.Strict().EncodeToString(bytes)

	// URL-encode Base64 string.
	b64String = strings.Replace(b64String, "+", "-", -1) // 62nd char of encoding
	b64String = strings.Replace(b64String, "/", "_", -1) // 63rd char of encoding

	return b64String
}
