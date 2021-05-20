package session

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gigamono/gigamono/pkg/configs"
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gin-gonic/gin"
)

// SetSessionCookie sets a session cookie.
// Sepcial attention to cookie attributes are important to keep sessions secure.
func SetSessionCookie(ctx *gin.Context, app *inits.App, key string, value string) error {
	if !strings.HasPrefix(key, "__Host") {
		return errors.New("session cookie must be prefixed with \"__Host\"")
	}

	secureAttr := false

	// Add additional cookie flags if environment is production.
	if app.Config.Environment == configs.Production {
		secureAttr = true
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     key,                  // __Host prefix to prevent subdomain CSRF hijacking.
		Value:    value,                // Cookie value
		Path:     "/",                  // Needed in combo with "__Host" prefix
		MaxAge:   604800,               // Cookie expires after a week.
		HttpOnly: true,                 // Cookie to be inaccessible to client code.
		Secure:   secureAttr,           // In production, only allow connection via HTTPS
		SameSite: http.SameSiteLaxMode, // Prevents certain CSRF-prone cross-origin requests.
	})

	return nil
}
