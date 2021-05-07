package logs

import (
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

// Println logs a message.
// This is just to provide a singe source of truth for everyting logging.
func Println(v ...interface{}) {
	log.Println(v...)
}

// Printf logs a message.
// This is just to provide a singe source of truth for everyting logging.
func Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
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
