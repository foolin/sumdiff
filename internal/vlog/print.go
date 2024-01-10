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
	write(fmt.Sprint(a...))
}

func write(msg string) {
	if !verbose {
		return
	}
	writeWith(printer, msg)
}

func writeWith(writer io.Writer, msg string) {
	msg = fmt.Sprintf("%v %v\n", time.Now().Format("2006-01-02 15:04:05.000"), msg)
	_, _ = writer.Write([]byte(msg))
}

func Printf(format string, a ...any) {
	write(fmt.Sprintf(format, a...))
}

func ExitWithCode(code int, msg string, a ...any) {
	writeWith(os.Stderr, fmt.Sprintf(msg, a...))
	os.Exit(code)
}

func Exit(msg string, a ...any) {
	ExitWithCode(1, msg, a...)
}
