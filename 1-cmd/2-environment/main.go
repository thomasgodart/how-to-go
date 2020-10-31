package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strconv"

	dotenv "github.com/joho/godotenv"
)

// --- initialization of the app with an env file,       ---
// --- then overwriting parameters with the command line ---

// we can put multiple "init" functions in one go file, like that:

// app init

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
}

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

// default values for command line parameters

var Word string = "foo"
var Numb int = 42
var Bool bool = true

func init() {

	// command line parameters setup

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
