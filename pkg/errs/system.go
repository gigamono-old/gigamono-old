package errs

import (
	"fmt"
	"runtime"

	"github.com/gigamono/gigamono/pkg/logs"
)

// SystemError represents an internal system error, typically invoked by a panic that will be handled by a panic handler.
// It wraps actual error to add context information to error.
type SystemError struct {
	ClientContextMessage string        // A contextual message for the client.
	ServerContextMessage string        // A contextual message for the server.
	FunctionFrame        runtime.Frame // The function frame that invoked a panic.
	ActualError          error         // The wrapped error.
}

// NewSystemError creates a new SystemError instance.
func NewSystemError(clientContextMessage string, serverContextMessage string, actualError error) SystemError {
	// Get the frame info of the routine that calls this function.
	frame := logs.CallerParentFrame()

	return SystemError{
		ClientContextMessage: clientContextMessage,
		ServerContextMessage: serverContextMessage,
		ActualError:          actualError,
		FunctionFrame:        frame,
	}
}

func (err *SystemError) Error() string {
	fileName := err.FunctionFrame.File
	funcName := err.FunctionFrame.Function
	lineInfo := err.FunctionFrame.Line
	return fmt.Sprintf(
		"%v: %v:\n:: %v:\n:: %v: %v\n",
		err.ServerContextMessage,
		err.ActualError,
		funcName,
		fileName,
		lineInfo,
	)
}

func (err *SystemError) Unwrap() error {
	return err.ActualError
}
