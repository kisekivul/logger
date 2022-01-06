package logger

import (
	"log"

	"github.com/sirupsen/logrus"
)

// LogHook log hook object
type LogHook struct{}

// Levels accept levels
func (lh *LogHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
		logrus.TraceLevel,
	}
}

// Fire fire hook
func (lh *LogHook) Fire(entry *logrus.Entry) error {
	log.Println(entry.Message)
	return nil
}
