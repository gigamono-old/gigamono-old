package middleware

import (
	"context"
	"fmt"

	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/messages"
	"github.com/gigamono/gigamono/pkg/services/rest/response"
	"github.com/gigamono/gigamono/pkg/services/session"
	"github.com/gin-gonic/gin"
)

// ClaimContextKey reepresents the key of JWT claim value stored in a context.
type ClaimContextKey struct {}

// Authenticate authenticates the user session.
func Authenticate(app *inits.App) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get public key.
		publicKey, err := app.Secrets.Get("AUTH_PUBLIC_KEY")
		if err != nil {
			panic(errs.NewSystemError(
				messages.Error["authenticate-user"].(string),
				"trying to retrieve public key from secrets manager",
				err,
			))
		}

		// Verify session tokens.
		claims, err := session.VerifySessionTokens(ctx, []byte(publicKey))
		if err != nil {
			switch err.(type) {
			case errs.ClientError:
				fmt.Print("Hello there. Failed to Authenticate")
				response.AuthenticationErrors(ctx, err.Error())
				ctx.Abort()
				return
			default:
				panic(errs.NewSystemError(
					messages.Error["signin"].(string),
					"verifying user session",
					err,
				))
			}
		}

		// Store claims in a new context.
		newCtx := context.WithValue(ctx.Request.Context(), ClaimContextKey{}, *claims)
		fmt.Println("Authenticated successfully", claims)

		// Replace request context.
		ctx.Request.WithContext(newCtx)
	}
}
