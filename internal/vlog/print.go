package vlog

import (
	"fmt"
	"io"
	"os"
	"time"
)

var printer io.Writer = os.Stdout

var verbose bool = false

func SetVerbose(v bool) {
	verbose = v
}

func Print(a ...any) {
	if !verbose {
		return
	}
	write(fmt.Sprint(a...))
}

func write(msg string) {
	msg = fmt.Sprintf("%v %v\n", time.Now().Format("2006-01-02 15:04:05"), msg)
	_, _ = printer.Write([]byte(msg))
}

func Printf(format string, a ...any) {
	if !verbose {
		return
	}
	write(fmt.Sprintf(format, a...))
}

func ExitWithCode(code int, msg string, a ...any) {
	Printf(msg, a...)
	os.Exit(code)
}

func Exit(code int, msg string, a ...any) {
	Exit(1, msg, a...)
}
