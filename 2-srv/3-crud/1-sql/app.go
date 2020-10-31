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

	// it is useful when the app is run as a systemd service, because the
	// current directory for a service is "/"

	// /!\ when this is used, then we can't "go run *.go" anymore,
	// because in that case golang compiles in /tmp/... and runs there

	executable, err := os.Executable()
	if err != nil {
		error := fmt.Sprintf("os.Executable() error: %s", err)
		panic(error)
	}
	err = os.Chdir(path.Dir(executable))
	if err != nil {
		error := fmt.Sprintf("os.Chdir() error: %s", err)
		panic(error)
	}
	dirRoot = path.Dir(executable)
}
