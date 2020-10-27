package main

import (
	"fmt"
	"os"
	"path"
)

// app init

var dirRoot string

func init() {

	// sets the current directory to where the app is located

	// /!\ when this is used, then we can't "go run *.go" anymore,
	// because in that case golang compiles in /tmp/... and runs there

	executable, err := os.Executable()
	if err != nil {
		panic(fmt.Sprintf("os.Executable() error: %s", err))
	}
	err = os.Chdir(path.Dir(executable))
	if err != nil {
		panic(fmt.Sprintf("os.Chdir() error: %s", err))
	}
	dirRoot = path.Dir(executable)
}
