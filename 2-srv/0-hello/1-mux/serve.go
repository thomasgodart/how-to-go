package main

import (
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {

	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message + "\n"

	w.Write([]byte(message))
}

func init() {

	Mux.HandleFunc("/", sayHello)
}
