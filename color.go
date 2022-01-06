package logger

import (
	"github.com/sirupsen/logrus"
)

var (
	Green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	White   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	Yellow  = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	Red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	Blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	Magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	Cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	Reset   = string([]byte{27, 91, 48, 109})
)

func ColorForMessage(level logrus.Level) string {
	switch level {
	case logrus.TraceLevel:
		return Cyan
	case logrus.DebugLevel:
		return Magenta
	case logrus.InfoLevel:
		return Green
	case logrus.WarnLevel:
		return Yellow
	case logrus.ErrorLevel:
		return Red
	default:
		return Reset
	}
}

func ColorForStatus(code int) string {
	switch {
	case code >= 200 && code < 300:
		return Green
	case code >= 300 && code < 400:
		return White
	case code >= 400 && code < 500:
		return Yellow
	default:
		return Red
	}
}

func ColorForMethod(method string) string {
	switch method {
	case "GET":
		return Blue
	case "POST":
		return Cyan
	case "PUT":
		return Yellow
	case "DELETE":
		return Red
	case "PATCH":
		return Green
	case "HEAD":
		return Magenta
	case "OPTIONS":
		return White
	default:
		return Reset
	}
}
