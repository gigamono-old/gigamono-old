package middleware

import (
	"context"

	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gigamono/gigamono/pkg/database/models/resource"
	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/messages"
	"github.com/gigamono/gigamono/pkg/security"
	"github.com/gigamono/gigamono/pkg/services/rest/response"
	"github.com/gigamono/gigamono/pkg/services/session"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// SessionData represnts sessionuser data.
type SessionData struct {
	UserID uuid.UUID
	Claims security.Claims
}

// SessionDataKey reepresents the key of JWT claim value stored in a context.
var SessionDataKey = struct{ key string }{key: "SessionDataKey"}

// AuthenticateCreateUser authenticates the user session and adds user if user does not exist.
func AuthenticateCreateUser(app *inits.App) gin.HandlerFunc {
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
				response.AuthenticationErrors(ctx, err.Error())
				ctx.Abort()
				return
			default:
				panic(errs.NewSystemError(
					messages.Error["authenticate-user"].(string),
					"verifying user session",
					err,
				))
			}
		}

		// Create session data.
		userID, err := uuid.FromString(claims.Subject)
		if err != nil {
			panic(errs.NewSystemError(
				messages.Error["authenticate-user"].(string),
				"converting user id from string to uuid",
				err,
			))
		}

		sessionData := SessionData{
			UserID: userID,
			Claims: *claims,
		}

		// Add new user if user does not already exist.
		user := resource.User{Base: models.Base{ID: sessionData.UserID}}
		if err = user.CreateIfNotExist(&app.DB); err != nil {
			panic(errs.NewSystemError(
				messages.Error["authenticate-user"].(string),
				"trying to create user if not exist",
				err,
			))
		}

		// Store claims in a new context.
		newCtx := context.WithValue(ctx.Request.Context(), SessionDataKey, sessionData)

		// Replace request context.
		ctx.Request = ctx.Request.WithContext(newCtx)
	}
}
