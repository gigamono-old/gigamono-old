package logs

import (
	"runtime"
)

// CallerFrame gets stack information of the caller.
func CallerFrame() runtime.Frame {
	programCounter := make([]uintptr, 15)
	callersCount := runtime.Callers(2, programCounter)
	frames := runtime.CallersFrames(programCounter[:callersCount])
	frame, _ := frames.Next()
	return frame
}

// CallerParentFrame gets stack information of the caller of the caller.
func CallerParentFrame() runtime.Frame {
	programCounter := make([]uintptr, 15)
	callersCount := runtime.Callers(2, programCounter)
	frames := runtime.CallersFrames(programCounter[:callersCount])
	frames.Next()
	frame, _ := frames.Next()
	return frame
}
