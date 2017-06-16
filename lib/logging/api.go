package logging

import (
	"golang.org/x/net/context"
)

// Logger offers an environment agnostic interface for logging
type Logger interface {
	Debug(c context.Context, format string, args ...interface{})
	Info(c context.Context, format string, args ...interface{})
	Warning(c context.Context, format string, args ...interface{})
	Error(c context.Context, format string, args ...interface{})
	Critical(c context.Context, format string, args ...interface{})
}

type loggerFactory func() Logger

// New is factory that return environment specific version of logger
var New loggerFactory
