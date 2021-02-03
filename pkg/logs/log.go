package logs

import (
	"os"
	"log"
	"fmt"

	"github.com/sageflow/sageutils/pkg/files"
)

// SetStatusLogFile sets the file where status logs go.
func SetStatusLogFile() {
	file, err := OpenLogFile("status.log")
	if err != nil {
		log.Printf("Cannot open or create 'logs/status.log' file: %v\nFalling back to stdout/stderr\n", err)
	}
	log.SetPrefix("\n")
	log.SetOutput(file)
}

// OpenLogFile opens/creates log file.
func OpenLogFile(dest string) (*os.File, error) {
	return files.OpenFile("logs/"+dest, true)
}

// FmtPrintln calls log and fmt Println function.
func FmtPrintln(v ...interface{}) {
	fmt.Println(v...)
	log.Println(v...)
}

// FmtPrintf calls log and fmt Printf function.
func FmtPrintf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
	log.Printf(format, v...)
}
