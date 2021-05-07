package middleware

import (
	"context"
	"errors"

	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/gigamono/gigamono/pkg/logs"
	"github.com/gigamono/gigamono/pkg/messages"
)

// PanicHandler logs all panics from graphql handler.
// Sec: DON'T log requests because of sensitive data.
func PanicHandler(_ context.Context, err interface{}) error {
	newErr := err

	// Check if err is a SystemError
	if systemErr, ok := err.(*errs.SystemError); ok {
		newErr = systemErr
	}

	// Log error.
	logs.FmtPrintln(newErr)

	// Send  system error to client.
	return errors.New(messages.Error["internal-system-error"].(string))
}
