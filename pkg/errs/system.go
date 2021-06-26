package errs

import (
	"fmt"

	"github.com/gigamono/gigamono/pkg/logs"
)

// SystemError represents an internal system error, typically invoked by a panic that will be handled by a panic handler.
// It wraps actual error to add context information to error.
type SystemError struct {
	ClientContextMessage string       // A contextual message for the client.
	ServerContextMessage string       // A contextual message for the server.
	FunctionInfo         FunctionInfo // Info about the function that invoked a panic.
	ActualError          error        // The wrapped error.
}

// FunctionInfo holds information about a function.
//
// Cannot store runtime.Frame directly because Gin has trouble printing it.
type FunctionInfo struct {
	Name string
	File string
	Line int
}

// NewSystemError creates a new SystemError instance.
func NewSystemError(clientContextMessage string, serverContextMessage string, actualError error) SystemError {
	// Get the frame info of the routine that calls this function.
	// Cannot store runtime.Frame directly because Gin has trouble printing it.
	frame := logs.CallerParentFrame()
	functionInfo := FunctionInfo{
		Name: frame.Function,
		File: frame.File,
		Line: frame.Line,
	}

	return SystemError{
		ClientContextMessage: clientContextMessage,
		ServerContextMessage: serverContextMessage,
		ActualError:          actualError,
		FunctionInfo:         functionInfo,
	}
}

func (err SystemError) Error() string {
	return fmt.Sprintf(
		"%v: %v:\n:: %v:\n:: %v:%v\n",
		err.ServerContextMessage,
		err.ActualError,
		err.FunctionInfo.Name,
		err.FunctionInfo.File,
		err.FunctionInfo.Line,
	)
}

func (err *SystemError) Unwrap() error {
	return err.ActualError
}
