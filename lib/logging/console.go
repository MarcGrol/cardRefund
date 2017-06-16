// +build !appengine

package logging

import (
	"log"

	"os"

	"golang.org/x/net/context"

	. "github.com/logrusorgru/aurora"
)

type consoleLogger struct {
	debug    *log.Logger
	info     *log.Logger
	warning  *log.Logger
	error    *log.Logger
	critical *log.Logger
}

func init() {
	New = newConsoleLogger
}

func newConsoleLogger() Logger {
	return &consoleLogger{
		debug:    log.New(os.Stderr, "duxxie", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile),
		info:     log.New(os.Stderr, "duxxie", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile),
		warning:  log.New(os.Stderr, "duxxie", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile),
		error:    log.New(os.Stderr, "duxxie", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile),
		critical: log.New(os.Stderr, "duxxie", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile),
	}
}

func (logger *consoleLogger) Debug(c context.Context, format string, args ...interface{}) {
	logger.debug.Printf(Sprintf(Green("DEBUG: "+format), args...))
}

func (logger *consoleLogger) Info(c context.Context, format string, args ...interface{}) {
	logger.info.Printf(Sprintf(Blue("INFO:  "+format), args...))
}

func (logger *consoleLogger) Warning(c context.Context, format string, args ...interface{}) {
	logger.warning.Printf(Sprintf(Magenta("WARN:  "+format), args...))
}

func (logger *consoleLogger) Error(c context.Context, format string, args ...interface{}) {
	logger.error.Printf(Sprintf(Red("ERROR: "+format), args...))
}

func (logger *consoleLogger) Critical(c context.Context, format string, args ...interface{}) {
	logger.critical.Printf(Sprintf(Red("FATAL: "+format), args...))
}
