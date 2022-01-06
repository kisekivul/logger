package logger

import (
	"path/filepath"
)

var (
	// Logger common logger instance
	logger *Logger
	file   = "logs/logs.log"
)

func Initialize(path string, mode Mode) {
	var (
		dir, file = filepath.Split(path)
	)
	// set logger
	logger = (&Logger{}).Initialize(dir, file, mode)
}

func initialize(path string, mode Mode) {
	var (
		dir, file = filepath.Split(path)
	)
	// set logger
	logger = (&Logger{}).Initialize(dir, file, mode)
}

func Println(args ...interface{}) {
	if logger == nil {
		initialize(file, TestMode)
	}
	logger.Println(args...)
}

func Infoln(args ...interface{}) {
	if logger == nil {
		initialize(file, TestMode)
	}
	logger.Infoln(args...)
}

func Debugln(args ...interface{}) {
	if logger == nil {
		initialize(file, TestMode)
	}
	logger.Debugln(args...)
}

func Warnln(args ...interface{}) {
	if logger == nil {
		initialize(file, TestMode)
	}
	logger.Warnln(args...)
}

func Errorln(args ...interface{}) {
	if logger == nil {
		initialize(file, TestMode)
	}
	logger.Errorln(args...)
}

func Fatalln(args ...interface{}) {
	if logger == nil {
		initialize(file, TestMode)
	}
	logger.Fatalln(args...)
}

func Panicln(args ...interface{}) {
	if logger == nil {
		initialize(file, TestMode)
	}
	logger.Panicln(args...)
}
