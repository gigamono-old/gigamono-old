package session

import (
	"fmt"

	"github.com/gigamono/gigamono/pkg/configs"
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gin-gonic/gin"
)

// SetSessionCookie sets a session cookie.
// The attributes are important to keep sessions secure.
// Max-Age represents how long until a cookie expires. This is set to a week in seconds.
// HttpOnly to make cookie inaccessible to client code.
// Path to specify allowed
// SameSite to prevent certain CSRF-prone cross-origin requests.
// Secure to enforce connection via HTTPS only.
func SetSessionCookie(ctx *gin.Context, app *inits.App, key string, value string) {
	attributes := "Max-Age=604800; HttpOnly; Path=/"

	// Add additional cookie flags if environment is production.
	if app.Config.Environment == configs.Production {
		attributes += "; SameSite=Lax; Secure"
	}

	ctx.Cookie(
		fmt.Sprintf("__Host%v:%v;%v", key, value, attributes), // __Host prefix to prevent subdomain CSRF hijacking.
	)
}
