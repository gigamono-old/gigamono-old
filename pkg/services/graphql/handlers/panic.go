package handlers

import (
	"context"
	"errors"

	"github.com/sageflow/sageflow/pkg/logs"
)

// PanicHandler logs all panics from graphql handler.
func PanicHandler(ctx context.Context, err interface{}) error {
	logs.FmtPrintln(err)
	return errors.New("internal system error")
}
