package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// database structure and init

type DBDoc struct {
	gorm.Model
	Name    string
	Content string
}

var db *gorm.DB

func init() {
	var err error

	db, err = gorm.Open(sqlite.Open("db/sqlite.db"), &gorm.Config{})
	if err != nil {
		error := fmt.Sprintf("gorm.Open(sqlite.Open(\"db/sqlite.db\") error: %s", err)
		panic(error)
	}

	db.AutoMigrate(&DBDoc{})
}
