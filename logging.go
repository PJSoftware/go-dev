package dev

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)

type Dev struct {
	logger         *log.Logger
	logFilename    string
	loggingEnabled bool
}

// InitLogging prepares a dev log file for use. However, by default logging is
// disabled until dev.EnableLogging() is called
func InitLogging(filePath string) *Dev {
	dev := &Dev{}

	folder, file := filepath.Split(filePath)
	if folder == "" {
		folder = tempFolder()
	}
	if path.Ext(file) == "" {
		file += ".log"
	}

	dir, err := filepath.Abs(folder)
	if err != nil {
		log.Fatalf("Error generating DevLog folder name: %v", err)
	}
	dev.initLogFile(dir, file)
	dev.loggingEnabled = false
	return dev
}

func (dev *Dev) EnableLogging() {
	dev.loggingEnabled = true
	fmt.Printf("Dev Logging enabled: consult '%s'\n", dev.logFilename)
}

func (dev *Dev) DisableLogging() {
	dev.loggingEnabled = false
	fmt.Printf("Dev Logging disabled\n")
}

func (dev *Dev) Print(v ...any) {
	if !dev.loggingEnabled { return }
	dev.logger.Print(v...)
}

func (dev *Dev) Printf(format string, v ...any) {
	if !dev.loggingEnabled { return }
	dev.logger.Printf(format, v...)
}

func tempFolder() string {
	temp := os.Getenv("TEMP")
	if temp == "" {
		temp = "."
	}
	exe := os.Args[0]
	_, ffn := filepath.Split(exe)
	fn := ffn[:len(ffn)-len(filepath.Ext(ffn))]
	return filepath.Join(temp, "Log", "Dev", fn)
}

// dev.initLogFile sets up logging to use the specified file
func (dev *Dev) initLogFile(logFolder, logFile string) {
	os.MkdirAll(logFolder, os.ModePerm)
	fp := filepath.Join(logFolder, logFile)
	dev.logFilename = fp

	w, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("dev log file not created: %v", err)
	}
	dev.logger = log.New(w, "[DEV] ", log.Ldate|log.Ltime|log.Llongfile)
}