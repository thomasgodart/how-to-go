package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// database directory init

func init() {

	dir := filepath.Join(dirRoot, "db")
	_, err := os.Stat(dir)
	if !os.IsNotExist(err) {
		return
	}

	err = os.Mkdir(dir, 0770)
	if err != nil {
		error := fmt.Sprintf("os.Mkdir(\"%s\", 0770) error: %s", dir, err)
		panic(error)
	}

	os.Chmod(dir, 0770) // because of https://github.com/golang/go/issues/25539#issuecomment-394472286
}

// database structure

type DBDoc struct {
	gorm.Model
	Name    string
	Content string
}

// database connection

var db *gorm.DB

func init() {
	var err error

	db, err = gorm.Open(sqlite.Open(filepath.Join(dirRoot, "db", "sqlite.db")), &gorm.Config{})
	if err != nil {
		error := fmt.Sprintf("gorm.Open(sqlite.Open(\"db/sqlite.db\") error: %s", err)
		panic(error)
	}

	db.AutoMigrate(&DBDoc{})
}
