package main

import (
	"fmt"
	"net"
	"net/http"

	router "github.com/gorilla/mux"
)

// different files are linked and compiled in alphabetical order,
// so the init functions are run in that order too

// app init

var Listen net.Listener

// the following can't be called "init" otherwise it's started when testing

func initListen() {
	var err error

	// this will open the listen on the network's interface

	Listen, err = net.Listen("tcp", "how-srv-simple:8080")
	if err != nil {
		error := fmt.Sprintf(`net.Listen("tcp", "how-srv-simple:8080") error: %s`, err)
		panic(error)
	}
}

var Mux *router.Router

func init() {

	// this will create a Gorilla Mux router

	Mux = router.NewRouter()
}

// main function

func main() {

	initListen()

	if err := http.Serve(Listen, Mux); err != nil {
		error := fmt.Sprintf("http.Serve(Listen, Mux) error: %s", err)
		panic(error)
	}
}
