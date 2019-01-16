package lazylog

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Loglevel loglevel
type Loglevel int

const (
	loglevelTrace Loglevel = iota
	loglevelDebug
	loglevelInfo
	loglevelError
	loglevelFatal
)

func (l *Loglevel) String() string {
	switch *l {
	case loglevelFatal:
		return "F"
	case loglevelError:
		return "E"
	case loglevelInfo:
		return "I"
	case loglevelDebug:
		return "D"
	case loglevelTrace:
		return "T"
	}
	return " "
}

// LazyLogger LazyLogger
type LazyLogger struct {
	loglevel Loglevel
}

func New(loglevel Loglevel) *LazyLogger {
	return &LazyLogger{loglevel: loglevel}
}

func NewFatal() *LazyLogger {
	return New(loglevelFatal)
}

func NewError() *LazyLogger {
	return New(loglevelError)
}

func NewInfo() *LazyLogger {
	return New(loglevelInfo)
}

func NewDebug() *LazyLogger {
	return New(loglevelDebug)
}

func NewTrace() *LazyLogger {
	return New(loglevelTrace)
}

func (l *LazyLogger) SetLoglevel(loglevel Loglevel) {
	l.loglevel = loglevel
}

func (l *LazyLogger) GetLoglevel() Loglevel {
	return l.loglevel
}

func (l *LazyLogger) Fatal(v ...interface{}) {
	Fatal(v...)
	os.Exit(1)
}

func (l *LazyLogger) Fatalf(format string, v ...interface{}) {
	Fatalf(format, v...)
	os.Exit(1)
}

func (l *LazyLogger) Error(v ...interface{}) {
	if l.loglevel <= loglevelError {
		Error(v...)
	}
}

func (l *LazyLogger) Errof(format string, v ...interface{}) {
	if l.loglevel <= loglevelError {
		Errof(format, v...)
	}
}

func (l *LazyLogger) Info(v ...interface{}) {
	if l.loglevel <= loglevelInfo {
		Info(v...)
	}
}

func (l *LazyLogger) Infof(format string, v ...interface{}) {
	if l.loglevel <= loglevelInfo {
		Infof(format, v...)
	}
}

func (l *LazyLogger) Debug(v ...interface{}) {
	if l.loglevel <= loglevelDebug {
		Debug(v...)
	}
}

func (l *LazyLogger) Debugf(format string, v ...interface{}) {
	if l.loglevel <= loglevelDebug {
		Debugf(format, v...)
	}
}

func (l *LazyLogger) Trace(v ...interface{}) {
	if l.loglevel > loglevelTrace {
		Trace(v...)
	}
}

func (l *LazyLogger) Tracef(format string, v ...interface{}) {
	if l.loglevel <= loglevelTrace {
		Tracef(format, v...)
	}
}

func Fatal(v ...interface{}) {
	writeln(os.Stderr, loglevelFatal, fmt.Sprint(v...))
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	writeln(os.Stderr, loglevelFatal, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Error(v ...interface{}) {
	writeln(os.Stderr, loglevelError, fmt.Sprint(v...))
}

func Errof(format string, v ...interface{}) {
	writeln(os.Stderr, loglevelError, fmt.Sprintf(format, v...))
}

func Info(v ...interface{}) {
	writeln(os.Stdout, loglevelInfo, fmt.Sprint(v...))
}

func Infof(format string, v ...interface{}) {
	writeln(os.Stdout, loglevelInfo, fmt.Sprintf(format, v...))
}

func Debug(v ...interface{}) {
	writeln(os.Stdout, loglevelDebug, fmt.Sprint(v...))
}

func Debugf(format string, v ...interface{}) {
	writeln(os.Stdout, loglevelDebug, fmt.Sprintf(format, v...))
}

func Trace(v ...interface{}) {
	writeln(os.Stdout, loglevelTrace, fmt.Sprint(v...))
}

func Tracef(format string, v ...interface{}) {
	writeln(os.Stdout, loglevelTrace, fmt.Sprintf(format, v...))
}

func writeln(w io.Writer, loglevel Loglevel, message string) {
	datetime := time.Now().Format("2006/1/2 15:04:05")
	fmt.Fprintln(w, datetime, loglevel.String(), message)
}
