package logger

import (
	"fmt"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// FatalStack like panic except it displays only the current goroutine's stack.
func (l *Logger) FatalStack(format string, v ...interface{}) {
	l.Logger.Fatal(fmt.Sprintf(format, v...) + "\n" + string(debug.Stack()))
}

// Data data list
func (l *Logger) Data(v ...interface{}) {
	l.Logger.Println(v...)
}

// Param param list
func (l *Logger) Param(length int, msg, param string) {
	l.Logger.Printf("%s%-"+strconv.Itoa(length)+"s %s --> %s%s%s\n", Magenta, msg, Reset, Cyan, param, Reset)
}

// Request request list
func (l *Logger) Request(date time.Time, method, url, ip string, latency time.Duration) {
	l.Logger.Printf("%s[%-7s]%s --> %s[%-30s]%s | %s[%-15s]%s | %s", ColorForMethod(method), method, Reset, ColorForMessage(logrus.InfoLevel), l.Abbreviate(url, 20), Reset, White, ip, Reset, latency)
}

// Abbreviate abbreviate sting
func (l *Logger) Abbreviate(src string, limit int) string {
	var (
		tar string
	)
	tar = src

	if len(src) > limit {
		tar = strings.Join([]string{string([]rune(src)[:limit]), "..."}, "")
	}
	return tar
}
