package session

import (
	"encoding/base64"
	"errors"
	"regexp"
	"strings"

	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/gin-gonic/gin"
)

// SetCustomHeader sets a custom header.
// Key must be prefixed with X-.
func SetCustomHeader(ctx *gin.Context, key string, value string) error {
	if !strings.HasPrefix(key, "X-") {
		return errors.New("custom header must be prefixed with \"X-\"")
	}

	ctx.Header(key, value)

	return nil
}

// GetBasicAuthCreds gets basic auth credentials from the header.
func GetBasicAuthCreds(ctx *gin.Context) (string, string, error) {
	// Get Authorization header value.
	value := ctx.GetHeader("Authorization")

	// Create prefix pattern.
	pattern, err := regexp.Compile("^Basic\\s+")
	if err != nil {
		return "", "", err
	}

	// Authorization header value must start with /Basic\s+/.
	if pattern.FindString(value) == "" {
		return "", "", errs.ClientError{
			Message: "Authorization header value must be prefixed with \"Basic \"",
		}
	}

	// Remove prefix pattern from value.
	credsString := pattern.ReplaceAllString(value, "")

	// Decode credentials value.
	credsBytes, err := base64.StdEncoding.DecodeString(credsString)
	if err != nil {
		return "", "", errs.ClientError{
			Message: "unable to decode base64 value",
		}
	}

	// Split value at semi-colon separator and get user-id (email) and password.
	// NOTE: email cannot contain semi-colon.
	creds := strings.SplitN(string(credsBytes), ":", 2)
	if len(creds) == 2 {
		return creds[0], creds[1], nil
	}

	return "", "", errs.ClientError{
		Message: "unable to parse Authorization header value",
	}
}
