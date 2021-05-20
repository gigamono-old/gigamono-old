package response

import (
	"net/http"

	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/gigamono/gigamono/pkg/messages"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// BindErrors sets binding error response. Returns 400.
//
// Gin form binding returns different types of errors. (ValidationErrors, NumError, etc.)
//
// https://github.com/gin-gonic/gin/issues/1093
func BindErrors(ctx *gin.Context, err error) {
	var clientErrors []errs.ClientError

	switch err.(type) {
	case validator.ValidationErrors: // Handle validation errors.
		for _, fieldErr := range err.(validator.ValidationErrors) {
			clientErrors = append(clientErrors, errs.ClientError{
				Path:    []string{ctx.FullPath(), fieldErr.Field()},
				Message: fieldErr.Error(),
				Code:    errs.InputValidationError,
				Type:    errs.URLEncodedForm,
			})
		}
	default: // Other types of errors.
		clientErrors = append(clientErrors, errs.ClientError{
			Path:    []string{ctx.FullPath()},
			Message: messages.Error["validation"].(messages.Func)("form"),
			Code:    errs.InputValidationError,
			Type:    errs.URLEncodedForm,
		})
	}

	// Send status.
	ctx.JSON(
		http.StatusBadRequest,
		Response{
			Errors: clientErrors,
		},
	)
}

// FormErrors sets form error response. Returns 400.
func FormErrors(ctx *gin.Context, code errs.MainErrorCode, message string) {
	ctx.JSON(
		http.StatusBadRequest,
		Response{
			Errors: []errs.ClientError{{
				Path:    []string{ctx.FullPath()},
				Message: message,
				Code:    code,
				Type:    errs.URLEncodedForm,
			}},
		},
	)
}

// BadRequestErrors sets response for invalid or bad request like validation errors. Returns 400.
func BadRequestErrors(ctx *gin.Context, err *errs.ClientError) {
	ctx.JSON(
		http.StatusBadRequest,
		Response{
			Errors: []errs.ClientError{*err},
		},
	)
}

// BasicAuthErrors sets basic auth error response. Returns 401.
func BasicAuthErrors(ctx *gin.Context, message string) {
	// Set a WWW-Authenticate header.
	ctx.Header("WWW-Authenticate", "Basic realm=\"PRODUCTION_RESOURCE_SERVER\", charset=\"UTF-8\"")

	// Return a 401 response.
	ctx.JSON(
		http.StatusUnauthorized,
		Response{
			Errors: []errs.ClientError{{
				Path:    []string{ctx.FullPath(), "Authorization"},
				Message: message,
				Code:    errs.InvalidBasicAuth,
				Type:    errs.Header,
			}},
		},
	)
}

// Success sets a success response. Returns 200.
func Success(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(
		http.StatusOK,
		Response{
			Message: message,
			Data:    data,
		},
	)
}
