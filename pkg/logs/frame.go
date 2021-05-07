package logs

import (
	"runtime"
)

// Frame gets stack information of the current function frame.
func Frame() runtime.Frame {
	programCounter := make([]uintptr, 15)
	callersCount := runtime.Callers(2, programCounter)
	frames := runtime.CallersFrames(programCounter[:callersCount])
	frame, _ := frames.Next()
	return frame
}
