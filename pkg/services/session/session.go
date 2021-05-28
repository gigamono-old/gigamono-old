package session

import (
	"fmt"

	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/messages"
	"github.com/gigamono/gigamono/pkg/security"
	"github.com/gin-gonic/gin"
)

// AttachPreSessionTokens sets pre-session CSRF ID cookie and custom header.
func AttachPreSessionTokens(ctx *gin.Context, app *inits.App, csrfID string, csrfJWT string) error {
	if err := SetSessionCookie(ctx, app, PreSessionAccessTokenCookie, csrfJWT); err != nil {
		return err
	}

	return SetCustomHeader(ctx, PreSessionCSRFHeader, csrfID)
}

// AttachSessionTokens sets session cookies and CSRF ID custom header.
func AttachSessionTokens(ctx *gin.Context, app *inits.App, csrfID string, accessToken string, refreshToken string) error {
	if err := SetSessionCookie(ctx, app, SessionAccessTokenCookie, accessToken); err != nil {
		return err
	}

	if err := SetSessionCookie(ctx, app, SessionRefreshTokenCookie, refreshToken); err != nil {
		return err
	}

	return SetCustomHeader(ctx, SessionCSRFHeader, csrfID)
}

// VerifyPreSessionCredentials verifies an existing pre-session CSRF ID and validates access token.
func VerifyPreSessionCredentials(ctx *gin.Context, publicKey []byte) error {
	// Fetch pre-session access token.
	accessToken, err := ctx.Cookie(PreSessionAccessTokenCookie)
	if err != nil {
		return errs.ClientError{
			Path:    []string{ctx.FullPath()},
			Message: messages.Error["pre-session-access-token-cookie"].(string),
			Code:    errs.PreSessionValidationError,
			Type:    errs.Cookie,
		}
	}

	// Fetch pre-session header.
	plaintextCSRFID := ctx.GetHeader(PreSessionCSRFHeader)

	// Decode and verify access token.
	payload, err := security.DecodeAndVerifySignedJWT(accessToken, publicKey)
	if err != nil {
		switch err.(type) {
		case errs.TamperError:
			return errs.ClientError{
				Path:    []string{ctx.FullPath()},
				Message: messages.Error["pre-session-csrf-tamper"].(string),
				Code:    errs.PreSessionValidationError,
				Type:    errs.Cookie,
			}
		case errs.ExpirationError:
			return errs.ClientError{
				Path:    []string{ctx.FullPath(), PreSessionAccessTokenCookie},
				Message: messages.Error["pre-session-csrf-expired"].(string),
				Code:    errs.PreSessionValidationError,
				Type:    errs.Cookie,
			}
		}
		return err
	}

	// Make sure token's action has the right value.
	if payload.Action != security.PreSession {
		return errs.ClientError{
			Path:    []string{ctx.FullPath()},
			Message: messages.Error["pre-session-access-token-tamper"].(string),
			Code:    errs.SessionValidationError,
			Type:    errs.Cookie,
		}
	}

	// Compare plaintext CSRF ID and signed CSRF ID.
	err = security.VerifySignedCSRFID(plaintextCSRFID, payload.SignedCSRFID, publicKey)
	switch err.(type) {
	case errs.TamperError:
		return errs.ClientError{
			Path:    []string{ctx.FullPath()},
			Message: messages.Error["pre-session-csrf-tamper"].(string),
			Code:    errs.PreSessionValidationError,
			Type:    errs.Cookie,
		}
	}

	return err
}

// VerifySessionTokens authenticates session user and returns session's claim
func VerifySessionTokens(ctx *gin.Context, publicKey []byte) (*security.Claims, error) {
	// Fetch session access token.
	accessToken, err := ctx.Cookie(SessionAccessTokenCookie)
	if err != nil {
		return &security.Claims{}, errs.ClientError{
			Path:    []string{ctx.FullPath()},
			Message: messages.Error["session-access-token-cookie"].(string),
			Code:    errs.PreSessionValidationError,
			Type:    errs.Cookie,
		}
	}

	// Fetch pre-session header.
	plaintextCSRFID := ctx.GetHeader(SessionCSRFHeader)

	// Decode and verify access token.
	payload, err := security.DecodeAndVerifySignedJWT(accessToken, publicKey)
	if err != nil {
		switch err.(type) {
		case errs.TamperError:
			return &security.Claims{}, errs.ClientError{
				Path:    []string{ctx.FullPath()},
				Message: messages.Error["session-csrf-tamper"].(string),
				Code:    errs.SessionValidationError,
				Type:    errs.Cookie,
			}
		case errs.ExpirationError:
			return &security.Claims{}, errs.ClientError{
				Path:    []string{ctx.FullPath(), PreSessionAccessTokenCookie},
				Message: messages.Error["session-csrf-expired"].(string),
				Code:    errs.SessionValidationError,
				Type:    errs.Cookie,
			}
		}
		return &security.Claims{}, err
	}

	// Make sure token's action has the right value.
	if payload.Action != security.SessionAccess {
		return &security.Claims{}, errs.ClientError{
			Path:    []string{ctx.FullPath()},
			Message: messages.Error["session-access-token-tamper"].(string),
			Code:    errs.SessionValidationError,
			Type:    errs.Cookie,
		}
	}

	// Compare plaintext CSRF ID and signed CSRF ID.
	err = security.VerifySignedCSRFID(plaintextCSRFID, payload.SignedCSRFID, publicKey)
	switch err.(type) {
	case errs.TamperError:
		fmt.Println("VerifySignedCSRFID TamperError: reached here")
		return &security.Claims{}, errs.ClientError{
			Path:    []string{ctx.FullPath()},
			Message: messages.Error["session-csrf-tamper"].(string),
			Code:    errs.SessionValidationError,
			Type:    errs.Cookie,
		}
	}

	return payload, err
}
