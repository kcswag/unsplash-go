package logger

import (
	"io"
	"log"
	"os"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init()  {
	file, err := os.OpenFile("log.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Info = log.New(io.MultiWriter(file, os.Stdout),
		"INFO: ",
		log.Ldate|log.Ltime)
	Info.SetFlags(log.Lshortfile|log.LstdFlags)

	Warning = log.New(io.MultiWriter(file, os.Stdout),
		"WARNING: ",
		log.Ldate|log.Ltime)
	Warning.SetFlags(log.Lshortfile|log.LstdFlags)

	Error = log.New(io.MultiWriter(file, os.Stderr),
		"ERROR: ",
		log.Ldate|log.Ltime)
	Error.SetFlags(log.Lshortfile|log.LstdFlags)
}

