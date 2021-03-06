package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// data storage init

var dirData string

func init() {

	dirData = "db"

	dir := filepath.Join(dirRoot, dirData)
	_, err := os.Stat(dir)
	if !os.IsNotExist(err) {
		return
	}

	err = os.Mkdir(dir, 0770)
	if err != nil {
		error := fmt.Sprintf(`os.Mkdir("%s", 0770) error: %s`, dir, err)
		panic(error)
	}

	os.Chmod(dir, 0770) // because of https://github.com/golang/go/issues/25539#issuecomment-394472286
}
