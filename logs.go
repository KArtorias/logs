package logs

import (
	"context"
	"fmt"
	"log"
	"runtime"
)

// Info
func CtxInfo(ctx context.Context, format string, args ...interface{}) {
	flashLogFile()
	_, file, line, _ := runtime.Caller(1)
	var logId interface{}
	if ctx != nil {
		logId = ctx.Value(logger.logKey)
	} else {
		logId = ""
	}
	format = fmt.Sprintf("INFO %s:%d %v %s", file, line, logId, format)
	wg.Add(1)
	go func() {
		log.Printf(format, args...)
		wg.Done()
	}()
}

// Warn
func CtxWarn(ctx context.Context, format string, args ...interface{}) {
	flashLogFile()
	_, file, line, _ := runtime.Caller(1)
	var logId interface{}
	if ctx != nil {
		logId = ctx.Value(logger.logKey)
	} else {
		logId = ""
	}
	format = fmt.Sprintf("WARN %s:%d %v %s", file, line, logId, format)
	wg.Add(1)
	go func() {
		log.Printf(format, args...)
		wg.Done()
	}()
}

// Error
func CtxError(ctx context.Context, format string, args ...interface{}) {
	flashLogFile()
	_, file, line, _ := runtime.Caller(1)
	var logId interface{}
	if ctx != nil {
		logId = ctx.Value(logger.logKey)
	} else {
		logId = ""
	}
	format = fmt.Sprintf("ERROR %s:%d %v %s", file, line, logId, format)
	wg.Add(1)
	go func() {
		log.Printf(format, args...)
		wg.Done()
	}()
}

// Fatal
func CtxFatal(ctx context.Context, format string, args ...interface{}) {
	flashLogFile()
	_, file, line, _ := runtime.Caller(1)
	var logId interface{}
	if ctx != nil {
		logId = ctx.Value(logger.logKey)
	} else {
		logId = ""
	}
	format = fmt.Sprintf("FATAL %s:%d %v %s", file, line, logId, format)
	wg.Add(1)
	go func() {
		log.Printf(format, args...)
		wg.Done()
	}()
}

func Info(format string, args ...interface{}) {
	CtxInfo(nil, format, args...)
}

func Warn(format string, args ...interface{}) {
	CtxWarn(nil, format, args...)
}

func Error(format string, args ...interface{}) {
	CtxError(nil, format, args...)
}

func Fatal(format string, args ...interface{}) {
	CtxFatal(nil, format, args...)
}
