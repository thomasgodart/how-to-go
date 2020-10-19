package main

import (
	"html/template"
	"net/http"
	"strings"

	router "github.com/gorilla/mux"
)

type Obj map[string]interface{}

func init() {

	muxStatic(Mux, "/css/", "./html/css/")
	muxStatic(Mux, "/img/", "./html/img/")
	muxStatic(Mux, "/js/", "./html/js/")

	funcMap := template.FuncMap{
		"trim": strings.Trim,
	}

	tpl, err := template.New("glob").Funcs(funcMap).ParseGlob("./html/*.html")
	if err != nil {
		panic(err)
	}

	Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "-.html", Obj{})
	})

	Mux.HandleFunc("/aaa", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "aaa.html", Obj{})
	})

	Mux.HandleFunc("/bbb", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "bbb.html", Obj{})
	})
}

func muxStatic(mux *router.Router, relativePath, root string) {
	mux.PathPrefix(relativePath).Handler(http.StripPrefix(relativePath, http.FileServer(http.Dir(root))))
}
