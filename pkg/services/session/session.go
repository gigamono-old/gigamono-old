package session

import (
	"github.com/gigamono/gigamono/pkg/auth"
	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/messages"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// Tokens are used to identify session.
type Tokens struct {
	AccessToken  string
	RefreshToken string
}

// GetSessionUser authenticates and returns session's user id from the access token.
func GetSessionUser(tokens Tokens) (uuid.UUID, error) {
	// TODO: Sec: Auth.
	// Validate session (not pre-session) csrf token.
	// Validate session JWT tokens.
	// Return JWT claim subject. (not uuid).
	return uuid.FromString("4b523a9f-1be2-45ca-99fc-005f12581467")
}

// AttachPreSessionTokens sets pre-session CSRF ID cookie and custom header.
func AttachPreSessionTokens(ctx *gin.Context, app *inits.App, csrfID string, csrfJWT string) error {
	if err := SetSessionCookie(ctx, app, PreSessionAccessTokenCookie, csrfJWT); err != nil {
		return err
	}

	return SetCustomHeader(ctx, PreSessionCSRFHeader, csrfID)
}

// AttachSessionTokens sets session cookies and CSRF ID custom header.
func AttachSessionTokens(ctx *gin.Context, app *inits.App, csrfID string, accessToken string, refreshToken string) error {
	if err := SetSessionCookie(ctx, app, AccessTokenCookie, accessToken); err != nil {
		return err
	}

	if err := SetSessionCookie(ctx, app, RefreshTokenCookie, refreshToken); err != nil {
		return err
	}

	return SetCustomHeader(ctx, SessionCSRFHeader, csrfID)
}

// VerifyPreSessionCredentials verifies an existing pre-session CSRF ID and validates access token.
func VerifyPreSessionCredentials(ctx *gin.Context, publicKey []byte) error {
	// Fetch pre-session access token.
	accessToken, err := ctx.Cookie(PreSessionAccessTokenCookie)
	if err != nil {
		return &errs.ClientError{
			Path:    []string{ctx.FullPath()},
			Message: messages.Error["pre-session-access-token-cookie"].(string),
			Code:    errs.PreSessionValidationError,
			Type:    errs.Cookie,
		}
	}

	// Fetch pre-session header.
	plaintextCSRFID := ctx.GetHeader(PreSessionCSRFHeader)

	// Decode and verify access token.
	payload, err := auth.DecodeAndVerifySignedJWT(accessToken, publicKey)
	if err != nil {
		switch err.(type) {
		case errs.TamperError:
			return &errs.ClientError{
				Path:    []string{ctx.FullPath()},
				Message: messages.Error["pre-session-csrf-tamper"].(string),
				Code:    errs.PreSessionValidationError,
				Type:    errs.Cookie,
			}
		case errs.ExpirationError:
			return &errs.ClientError{
				Path:    []string{ctx.FullPath(), PreSessionAccessTokenCookie},
				Message: messages.Error["pre-session-csrf-expired"].(string),
				Code:    errs.PreSessionValidationError,
				Type:    errs.Cookie,
			}
		}
		return err
	}

	// Compare plaintext CSRF ID and signed CSRF ID.
	err = auth.VerifySignedCSRFID(plaintextCSRFID, payload.SignedCSRFID, publicKey)
	switch err.(type) {
	case errs.TamperError:
		return &errs.ClientError{
			Path:    []string{ctx.FullPath()},
			Message: messages.Error["pre-session-csrf-tamper"].(string),
			Code:    errs.PreSessionValidationError,
			Type:    errs.Cookie,
		}
	}

	return err
}
