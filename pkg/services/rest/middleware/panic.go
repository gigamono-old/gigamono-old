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
//
// Sec: DON'T log requests because of sensitive data.
func PanicHandler(ctx *gin.Context, err interface{}) {
	clientErrorMessage := messages.Error["internal"].(string)

	// Check if err is a SystemError and log error.
	if systemErr, ok := err.(errs.SystemError); ok {
		if systemErr.ClientContextMessage != "" {
			clientErrorMessage = systemErr.ClientContextMessage
		}
		logs.FmtPrintln(systemErr.Error())
	} else {
		logs.FmtPrintln(err)
	}

	// Send system error to client.
	ctx.JSON(
		http.StatusInternalServerError,
		response.Response{
			Errors: []errs.ClientError{{
				Path:            []string{ctx.FullPath()},
				Message:         clientErrorMessage,
				Code:            errs.InternalSystemError,
				Type:            errs.ClientErrorTypeNone,
			}},
		},
	)
}
