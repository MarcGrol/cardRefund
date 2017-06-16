// +build appengine

package logging

import (
	"golang.org/x/net/context"

	"google.golang.org/appengine/log"
)

type appengineLogger struct{}

func newAppengineLogger() Logger {
	return &appengineLogger{}
}

func (logger *appengineLogger) Debug(c context.Context, format string, args ...interface{}) {
	if c.Err() != nil {
		return
	}
	log.Debugf(c, format, args...)
}

func (logger *appengineLogger) Info(c context.Context, format string, args ...interface{}) {
	if c.Err() != nil {
		return
	}
	log.Infof(c, format, args...)
}

func (logger *appengineLogger) Warning(c context.Context, format string, args ...interface{}) {
	if c.Err() != nil {
		return
	}
	log.Warningf(c, format, args...)
}

func (logger *appengineLogger) Error(c context.Context, format string, args ...interface{}) {
	if c.Err() != nil {
		return
	}
	log.Errorf(c, format, args...)
}

func (logger *appengineLogger) Critical(c context.Context, format string, args ...interface{}) {
	if c.Err() != nil {
		return
	}
	log.Criticalf(c, format, args...)
}

func init() {
	New = newAppengineLogger
}
