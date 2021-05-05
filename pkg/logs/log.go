package logs

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/gigamono/gigamono/pkg/files"
)

// SetStatusLogFile sets the file where status logs go.
func SetStatusLogFile() {
	file, err := OpenOrCreateLogFile("status.log")
	if err != nil {
		log.Printf("Cannot open or create 'logs/status.log' file: %v\nFalling back to stdout/stderr\n", err)
	}
	log.SetPrefix("\n")
	log.SetOutput(file)
}

// OpenOrCreateLogFile opens/creates log file.
func OpenOrCreateLogFile(dest string) (*os.File, error) {
	return files.OpenOrCreateFile("logs/"+dest, true)
}

// FmtPrintln logs message and prints it to the console.
func FmtPrintln(v ...interface{}) {
	fmt.Println(v...)
	log.Println(v...)
}

// FmtPrintf logs message and prints it to the console.
func FmtPrintf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
	log.Printf(format, v...)
}

// Error logs message and returns an error message with the original error.
func Error(message string, err error) error {
	msg := message + ":"
	log.Println(msg, err)
	return errors.New(fmt.Sprintln(msg, err))
}

// NewError logs message and returns an error message but without the original error.
func NewError(message string, err error) error {
	log.Println(message+":", err)
	return fmt.Errorf(message)
}
