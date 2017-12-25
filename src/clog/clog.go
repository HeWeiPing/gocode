package clog

import (
	"fmt"
	"os"
	"runtime"
)

const (
	Black     = 30
	Red       = 31
	Green     = 32
	Yellow    = 33
	Blue      = 34
	Pink      = 35
	DeepGreen = 36
	White     = 37
)

func Clog(color int, format string, a ...interface{}) (n int, err error) {
	fmt.Printf("\033[;%d;1m", color)
	n, err = fmt.Fprintf(os.Stdout, format, a...)
	fmt.Printf("\033[0m\n")

	return n, err
}

//log out verbose
func Clogv(color int, format string, a ...interface{}) (n int, err error) {
	var ok bool
	var file string
	var line int

	_, file, line, ok = runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}

	fmt.Printf("file:\033[;35;1m%s:\033[0mline:\033[;33;1m%d\033[0m:_", file, line)
	fmt.Printf("\033[;%d;1m", color)
	n, err = fmt.Fprintf(os.Stdout, format, a...)
	fmt.Printf("\033[0m\n")

	return n, err
}
