package main

import (
	"flag"
	"fmt"
)

// --- command lines parameters parsing ---

// from https://gobyexample.com/command-line-flags

// default values

var wordPtr *string
var numbPtr *int
var boolPtr *bool

// app init

func init() {

	wordPtr = flag.String("word", "foo", "a string")
	numbPtr = flag.Int("numb", 42, "an int")
	boolPtr = flag.Bool("bool", false, "a bool")

	flag.Parse()
}

// main function

func main() {

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("bool:", *boolPtr)
}
