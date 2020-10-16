package main

import (
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {

	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message

	w.Write([]byte(message))
}

func init() {

	http.HandleFunc("/", sayHello)
}

func main() {

	// it's a good practice to name all your projects
	// in /etc/hosts and give them unique IPs, like:
	//
	// 127.0.0.1 localhost
	// 127.1.2.0 how-srv-hello
	// 127.1.2.1 how-srv-...
	//
	// then in a project, different services can each receive a different port

	if err := http.ListenAndServe("how-srv-hello:8080", nil); err != nil {
		panic(err)
	}
}
