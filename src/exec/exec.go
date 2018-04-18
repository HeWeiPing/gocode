package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := new(exec.Cmd)
	cmd.Path = "/bin/ls"
	cmd.Args = []string{"-la"}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Print(err)
	}
}
