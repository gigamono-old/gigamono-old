package session

import (
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gin-gonic/gin"
)

// AttachCSRFTokens sets pre-session CSRF token cookie and custom header.
func AttachCSRFTokens(ctx *gin.Context, app *inits.App, csrfToken string, signedCSRFToken string) {
	SetSessionCookie(ctx, app, "CSRFToken", csrfToken)
	SetCustomHeader(ctx, "Signed-CSRF-Token", signedCSRFToken)
}

// AttachPreSessionCSRFTokens sets pre-session CSRF token cookie and custom header.
func AttachPreSessionCSRFTokens(ctx *gin.Context, app *inits.App, csrfToken string, signedCSRFToken string) {
	SetSessionCookie(ctx, app, "PreSessionCSRFToken", csrfToken)
	SetCustomHeader(ctx, "Signed-Pre-Session-CSRF-Token", signedCSRFToken)
}
