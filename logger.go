package logger

import (
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

// Mode run mode of logger
type Mode string

const (
	// TestMode indicates mode is test.
	TestMode Mode = "test"
	// DebugMode indicates mode is debug.
	DebugMode Mode = "debug"
	// ReleaseMode indicates mode is release.
	ReleaseMode Mode = "release"
)

type Logger struct {
	*logrus.Logger
	mode Mode
	name string
}

// Initialize initialize logger
func (l *Logger) Initialize(dir, file string, mode Mode) *Logger {
	l.mode = mode
	//日志文件
	l.name = path.Join(dir, file)
	l.dir(dir)
	// init logger
	l.Logger = logrus.New()

	out, _ := rotatelogs.New(
		l.name+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(l.name),
		rotatelogs.WithMaxAge(time.Duration(24)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)

	switch mode {
	case TestMode:
		l.Logger.SetLevel(logrus.TraceLevel)
	case DebugMode:
		l.Logger.SetLevel(logrus.DebugLevel)
	case ReleaseMode:
		l.Logger.SetLevel(logrus.InfoLevel)
	}
	l.Logger.SetOutput(out)
	l.Logger.AddHook(&LogHook{})

	return l
}

func (l *Logger) dir(dir string) {
	//写入文件
	if f, err := os.Stat(dir); os.IsNotExist(err) || !f.IsDir() {
		os.MkdirAll(dir, 0755)
	}
}
