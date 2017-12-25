package clog_test

import (
	"clog"
	"testing"
)

func TestPrintlnRed(t *testing.T) {
	clog.Clog(clog.Red, "RRRRRRRRRRRRRRRRRRR a=%d b=%f", 5, 3.24)
}

func TestPrintlnGreen(t *testing.T) {
	clog.Clog(clog.Green, "GGGGGGGGGGGGGGGGGGG a=%d b=%f", 3, 3.24)
}

func TestPrintlnVerbose(t *testing.T) {
	clog.Clogv(clog.Green, "a=%d b=%f", 3, 3.24)
}
