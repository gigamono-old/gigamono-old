package middleware

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/gigamono/gigamono/pkg/logs"
	"github.com/gigamono/gigamono/pkg/messages"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// PanicHandler logs all panics from graphql handler.
//
// Sec: DON'T log requests because of sensitive data.
func PanicHandler(ctx context.Context, err interface{}) error {
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
	return &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: clientErrorMessage,
		Extensions: map[string]interface{}{
			"code": errs.InputValidationError,
		},
	}
}
