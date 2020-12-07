package main

import (
	"fmt"
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

	// this will register this handle to DefaultServeMux

	http.HandleFunc("/", sayHello)
}

func main() {

	// it's a good practice to name all your projects
	// in /etc/hosts and give them unique IPs, like:
	//
	// 127.0.0.1 localhost
	// 127.1.0.0 how-srv-hello
	// 127.1.1.0 how-srv-...
	//
	// then in a project, different services can each receive a different port

	// this will serve DefaultServeMux

	if err := http.ListenAndServe("how-srv-hello:8080", nil); err != nil {
		error := fmt.Sprintf(`http.ListenAndServe("how-srv-hello:8080", nil) error: %s`, err)
		panic(error)
	}
}
