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
	programCounters := make([]uintptr, 5)
	callersCount := runtime.Callers(3, programCounters)
	frames := *runtime.CallersFrames(programCounters[:callersCount])
	frame, _ := frames.Next()
	return frame
}
