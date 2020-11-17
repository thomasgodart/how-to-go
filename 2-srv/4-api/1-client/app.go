package main

import (
	"os"
)

// app init

var dirRoot string

func init() {
	var err error

	// gets the current directory

	dirRoot, err = os.Getwd()
	if err != nil {
		dirRoot = "."
	}
}
