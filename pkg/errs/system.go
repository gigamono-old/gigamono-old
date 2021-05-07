package errs

import (
	"fmt"
	"runtime"
)

// SystemError represents an internal system error, typically invoked by a panic that will be handled by a panic handler.
// It wraps actual error to add context information to error.
type SystemError struct {
	Context string        // A contextual message.
	Frame   runtime.Frame // The function frame that invoked a panic.
	Actual  error         /// The wrapped error.
}

// NewSystemError creates a new SystemError instance.
func NewSystemError(ctx string, actual error, frame runtime.Frame) SystemError {
	return SystemError{
		Context: ctx,
		Actual:  actual,
		Frame:   frame,
	}
}

func (err *SystemError) Error() string {
	fileName := err.Frame.File
	funcName := err.Frame.Function
	lineInfo := err.Frame.Line
	return fmt.Sprintf(
		"%v: %v: \n\t\t> %v: %v: %v\n",
		err.Context,
		err.Actual,
		fileName,
		lineInfo,
		funcName,
	)
}

func (err *SystemError) Unwrap() error {
	return err.Actual
}
