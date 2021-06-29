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

// SessionData represents sessionuser data.
type SessionData struct {
	User   *resource.User
	Claims *security.Claims
}

// SessionDataKey represents the key of JWT claim value stored in a context.
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

		// Convert userID
		userID, err := uuid.FromString(claims.Subject)
		if err != nil {
			panic(errs.NewSystemError(
				messages.Error["authenticate-user"].(string),
				"converting user id from string to uuid",
				err,
			))
		}

		// Add new user if user does not already exist.
		// TODO: Is there a way prevent this extra DB hit on every request? Maybe through adding extra info to claims in JWT.
		user := resource.User{Base: models.Base{ID: userID}}
		userExists, err := user.Exists(&app.DB)
		if err != nil {
			panic(errs.NewSystemError(
				messages.Error["authenticate-user"].(string),
				"checking if user exists",
				err,
			))
		}

		if !userExists {
			// Create user.
			if err := user.Create(&app.DB); err != nil {
				panic(errs.NewSystemError(
					messages.Error["authenticate-user"].(string),
					"trying to create user if not exist",
					err,
				))
			}

			// Create associated profile.
			profile := resource.Profile{UserID: user.ID, Email: &claims.Email}
			if err := profile.Create(&app.DB); err != nil {
				panic(errs.NewSystemError(
					messages.Error["authenticate-user"].(string),
					"trying to create user's profile",
					err,
				))
			}

			// Create associated preferences.
			preferences := resource.Preferences{UserID: user.ID}
			if err := preferences.Create(&app.DB); err != nil {
				panic(errs.NewSystemError(
					messages.Error["authenticate-user"].(string),
					"trying to create user's preferences",
					err,
				))
			}
		}

		// Populate SessionData.
		sessionData := SessionData{
			User:   &user,
			Claims: claims,
		}

		// Store claims in a new context.
		newCtx := context.WithValue(ctx.Request.Context(), SessionDataKey, sessionData)

		// Replace request context.
		ctx.Request = ctx.Request.WithContext(newCtx)
	}
}
