package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	dotenv "github.com/joho/godotenv"
)

// --- initialization of the app with an env file,    ---
// --- then overwriting parameters with command lines ---

// default values

var Word string = "foo"
var Numb int = 42
var Bool bool = true

// app init

func init() {

	// env file load and parsing

	// see https://github.com/joho/godotenv

	dotenv.Load()

	if envWord := os.Getenv("word"); envWord != "" {
		Word = envWord
	}
	if envNumb := os.Getenv("numb"); envNumb != "" {
		n, err := strconv.ParseInt(envNumb, 10, 0)
		if err != nil {
			panic(err)
		}
		Numb = int(n)
	}
	if envBool := os.Getenv("bool"); envBool != "" {
		b, err := strconv.ParseBool(envBool)
		if err != nil {
			panic(err)
		}
		Bool = b
	}
}

func init() { // we can put multiple "init" functions in one go file

	// command line parameters

	flag.StringVar(&Word, "word", Word, "a string")
	flag.IntVar(&Numb, "numb", Numb, "an int")
	flag.BoolVar(&Bool, "bool", Bool, "a bool")

	flag.Parse()
}

// main function

// entry point of the app AFTER all the init functions have run
// from top to bottom in each file, and files in alphabetical order

func main() {

	fmt.Println("word:", Word)
	fmt.Println("numb:", Numb)
	fmt.Println("bool:", Bool)
}
