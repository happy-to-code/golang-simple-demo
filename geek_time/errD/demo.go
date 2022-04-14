package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
	"sync/atomic"
)

type Error struct {
	err     error
	msg     string
	fullMsg string
	stackTrace
}

func Errf(err error, format string, args ...interface{}) *Error {
	e := &Error{
		err: err,
		msg: fmt.Sprintf(format, args),
	}
	e.fullMsg = e.getStackTrace()
	return e
}

func Err(err error, args ...interface{}) *Error {
	e := &Error{
		err: err,
		msg: fmt.Sprint(args),
	}
	e.fullMsg = e.getStackTrace()
	return e
}

type stackTrace struct {
	// stack info
	data    string
	callers []uintptr
}

func (err *Error) getStackTrace() string {
	if strings.TrimSpace(err.data) == "" {
		return err.genStackTrace(5)
	}
	return err.data
}

func (err *Error) StackTrace() string {
	return err.data
}

func (err *Error) genStackTrace(skip int) string {
	if config.isPrintStack == 1 {
		var buffer bytes.Buffer
		buffer.WriteString("StackTrace:\n")
		var st [64]uintptr
		n := runtime.Callers(skip, st[:])
		err.callers = st[:n]
		frames := runtime.CallersFrames(err.callers)
		for {
			frame, ok := frames.Next()
			if !ok {
				break
			}
			if !strings.Contains(frame.File, "runtime/") {
				buffer.WriteString(fmt.Sprintf("%s\n\t%s:%d\n",
					frame.Func.Name(), frame.File, frame.Line))
			}
		}
		err.data = buffer.String()
		return err.data
	}
	return ""
}

const (
	// PRINTSTACK print stack info
	PRINTSTACK = 1
)

// global error config object

var config *errConfig = &errConfig{}

// error config

type errConfig struct {
	isPrintStack uint32
}

// SetConfig set error config info
func SetConfig(conf byte) {
	if (conf & PRINTSTACK) != 0 {
		atomic.CompareAndSwapUint32(&config.isPrintStack, 0, 1)
	}
}
