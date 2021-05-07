package middleware

import (
	"net/http"

	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/gigamono/gigamono/pkg/logs"
	"github.com/gigamono/gigamono/pkg/messages"
	"github.com/gigamono/gigamono/pkg/services/rest/response"
	"github.com/gin-gonic/gin"
)

// PanicHandler logs all panics from rest endpoints.
// Sec: DON'T log requests because of sensitive data.
func PanicHandler(ctx *gin.Context, err interface{}) {
	newErr := err

	// Check if err is a SystemError
	if systemErr, ok := err.(*errs.SystemError); ok {
		newErr = systemErr
	}

	// Log error.
	logs.FmtPrintln(newErr)

	// Send system error to client.
	ctx.JSON(
		http.StatusInternalServerError,
		response.Response{
			Errors: []errs.ClientError{{
				Path:    []string{ctx.FullPath()},
				Message: messages.Error["internal-system-error"].(string),
				Code:    errs.InternalSystemError,
				Type:    errs.None,
			}},
		},
	)
}
