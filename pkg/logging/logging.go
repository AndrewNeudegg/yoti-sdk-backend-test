// Package logging wraps logrus imported logger for quick replacement.
package logging

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/logrusorgru/aurora"
	"github.com/sirupsen/logrus"
)

var (
	logger = logrus.New()
)

// formatMessage acts like middleware for the logger.
// It will insert additional data into the log messages.
func formatMessage(msg string) string {
	thisMsg := msg
	if config.DescribeCaller {
		file, line, funcName := getTrace()
		file = filepath.Base(file)
		funcName = filepath.Base(funcName)

		callerInfo := aurora.Gray(4, fmt.Sprintf("[%s:%s:%s]", file, line, funcName))
		thisMsg = fmt.Sprintf("%v %s", callerInfo, thisMsg)
	}
	return thisMsg
}

// Info will emit an info log to the output.
func Info(msg string) {
	logger.Info(formatMessage(msg))
}

// Error will emit an error log to the output, formatting an error to a string.
func Error(err error) {
	logger.Error(formatMessage(err.Error()))
}

// Errors will emit an error log to the output, from string.
func Errors(err string) {
	logger.Error(formatMessage(err))
}

// internalError may be called when configuring the logger.
func internalError(err error) {
	panic(err)
}

// getTrace will retrieve the caller from the stack trace.
func getTrace() (string, string, string) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(4, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return frame.File, strconv.Itoa(frame.Line), frame.Function
}
