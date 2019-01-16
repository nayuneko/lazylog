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
	LoglevelTrace Loglevel = iota
	LoglevelDebug
	LoglevelInfo
	LoglevelError
	LoglevelFatal
)

func (l *Loglevel) String() string {
	switch *l {
	case LoglevelFatal:
		return "F"
	case LoglevelError:
		return "E"
	case LoglevelInfo:
		return "I"
	case LoglevelDebug:
		return "D"
	case LoglevelTrace:
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
	return New(LoglevelFatal)
}

func NewError() *LazyLogger {
	return New(LoglevelError)
}

func NewInfo() *LazyLogger {
	return New(LoglevelInfo)
}

func NewDebug() *LazyLogger {
	return New(LoglevelDebug)
}

func NewTrace() *LazyLogger {
	return New(LoglevelTrace)
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
	if l.loglevel <= LoglevelError {
		Error(v...)
	}
}

func (l *LazyLogger) Errof(format string, v ...interface{}) {
	if l.loglevel <= LoglevelError {
		Errof(format, v...)
	}
}

func (l *LazyLogger) Info(v ...interface{}) {
	if l.loglevel <= LoglevelInfo {
		Info(v...)
	}
}

func (l *LazyLogger) Infof(format string, v ...interface{}) {
	if l.loglevel <= LoglevelInfo {
		Infof(format, v...)
	}
}

func (l *LazyLogger) Debug(v ...interface{}) {
	if l.loglevel <= LoglevelDebug {
		Debug(v...)
	}
}

func (l *LazyLogger) Debugf(format string, v ...interface{}) {
	if l.loglevel <= LoglevelDebug {
		Debugf(format, v...)
	}
}

func (l *LazyLogger) Trace(v ...interface{}) {
	if l.loglevel <= LoglevelTrace {
		Trace(v...)
	}
}

func (l *LazyLogger) Tracef(format string, v ...interface{}) {
	if l.loglevel <= LoglevelTrace {
		Tracef(format, v...)
	}
}

func Fatal(v ...interface{}) {
	writeln(os.Stderr, LoglevelFatal, fmt.Sprint(v...))
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	writeln(os.Stderr, LoglevelFatal, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Error(v ...interface{}) {
	writeln(os.Stderr, LoglevelError, fmt.Sprint(v...))
}

func Errof(format string, v ...interface{}) {
	writeln(os.Stderr, LoglevelError, fmt.Sprintf(format, v...))
}

func Info(v ...interface{}) {
	writeln(os.Stdout, LoglevelInfo, fmt.Sprint(v...))
}

func Infof(format string, v ...interface{}) {
	writeln(os.Stdout, LoglevelInfo, fmt.Sprintf(format, v...))
}

func Debug(v ...interface{}) {
	writeln(os.Stdout, LoglevelDebug, fmt.Sprint(v...))
}

func Debugf(format string, v ...interface{}) {
	writeln(os.Stdout, LoglevelDebug, fmt.Sprintf(format, v...))
}

func Trace(v ...interface{}) {
	writeln(os.Stdout, LoglevelTrace, fmt.Sprint(v...))
}

func Tracef(format string, v ...interface{}) {
	writeln(os.Stdout, LoglevelTrace, fmt.Sprintf(format, v...))
}

func writeln(w io.Writer, loglevel Loglevel, message string) {
	datetime := time.Now().Format("2006/01/02 15:04:05")
	fmt.Fprintln(w, datetime, loglevel.String(), message)
}
