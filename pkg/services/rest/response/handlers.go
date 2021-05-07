package response

import (
	"net/http"

	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/gigamono/gigamono/pkg/messages"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// SendFormBindErrors handles and sends binding errors.
// Gin form binding returns different types of errors. (ValidationErrors, NumError, etc.)
// https://github.com/gin-gonic/gin/issues/1093
func SendFormBindErrors(ctx *gin.Context, err error) {
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
			Message: messages.Error["input-validation"].(messages.Func)("form"),
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

// SendFormErrors simplifies sending form errors.
func SendFormErrors(ctx *gin.Context, code errs.ErrorCode, message string) {
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

// SendSuccess simplifies sending success messages.
func SendSuccess(ctx *gin.Context, message string, data map[string]interface{}) {
	ctx.JSON(
		http.StatusOK,
		Response{
			Message: message,
			Data:    data,
		},
	)
}
