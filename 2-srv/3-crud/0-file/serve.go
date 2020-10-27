package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	router "github.com/gorilla/mux"
)

type Obj map[string]interface{}

type Doc struct {
	URL     string
	Name    string
	Content string
}

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

	Mux.HandleFunc("/all", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {

			read := filepath.Join(dirRoot, "db")
			dirs, err := ioutil.ReadDir(read)
			if err != nil {
				error := fmt.Sprintf("can't read directory '%s'", read)
				muxError(error)
				url := path.Join("/")
				http.Redirect(w, r, url, http.StatusSeeOther)
				return
			}

			docs := make(map[string]Doc)

			for _, dir := range dirs {

				docs[dir.Name()] = Doc{
					URL:  path.Join("/", "get", dir.Name()),
					Name: dir.Name(),
				}
			}

			o := Obj{}
			o["Docs"] = docs
			tpl.ExecuteTemplate(w, "all.html", o)
		}

		if r.Method == "POST" {
			url := path.Join("/all")
			http.Redirect(w, r, url, http.StatusSeeOther)
		}
	})

	Mux.HandleFunc("/get/{name}", func(w http.ResponseWriter, r *http.Request) {
		route := router.Vars(r)

		if r.Method == "GET" {
			name := muxPath(route["name"])

			dir := filepath.Join(dirRoot, "db", name)
			_, err := os.Stat(dir)

			if os.IsNotExist(err) {
				w.WriteHeader(http.StatusNotFound)
				tpl.ExecuteTemplate(w, "404.html", Obj{})
				return
			}

			var doc Doc
			doc.Name = name

			if binary, err := ioutil.ReadFile(filepath.Join(dir, "content.html")); err == nil {
				doc.Content = string(binary)
			}

			o := Obj{}
			o["Doc"] = doc
			tpl.ExecuteTemplate(w, "get.html", o)
		}

		if r.Method == "POST" {
			url := path.Join("/get")
			http.Redirect(w, r, url, http.StatusSeeOther)
		}
	})

	Mux.HandleFunc("/set/{name}", func(w http.ResponseWriter, r *http.Request) {
		route := router.Vars(r)

		if r.Method == "GET" {
			name := muxPath(route["name"])

			dir := filepath.Join(dirRoot, "db", name)
			_, err := os.Stat(dir)

			if os.IsNotExist(err) {
				w.WriteHeader(http.StatusNotFound)
				tpl.ExecuteTemplate(w, "404.html", Obj{})
				return
			}

			var doc Doc
			doc.Name = name

			if binary, err := ioutil.ReadFile(filepath.Join(dir, "content.html")); err == nil {
				doc.Content = string(binary)
			}

			o := Obj{}
			o["Doc"] = doc
			tpl.ExecuteTemplate(w, "set.html", o)
		}

		if r.Method == "POST" {
			r.ParseForm()

			name := muxPath(r.FormValue("name"))
			content := r.FormValue("content")

			dir := filepath.Join(dirRoot, "db", name)
			_, err := os.Stat(dir)

			if os.IsNotExist(err) {
				error := fmt.Sprintf("document '%s' does not exist", name)
				muxError(error)
				url := path.Join("/set", name)
				http.Redirect(w, r, url, http.StatusSeeOther)
				return
			}

			f, err := os.Create(filepath.Join(dir, "content.html"))
			defer f.Close()

			if err != nil {
				url := path.Join("/set", name)
				http.Redirect(w, r, url, http.StatusSeeOther)
				return
			}

			many, err := fmt.Fprint(f, content)
			if err != nil || many != len(content) {
				// error
			} else {
				// success
			}

			url := path.Join("/get", name)
			http.Redirect(w, r, url, http.StatusSeeOther)
		}
	})

	Mux.HandleFunc("/del/{name}", func(w http.ResponseWriter, r *http.Request) {
		route := router.Vars(r)

		if r.Method == "GET" {
			name := muxPath(route["name"])

			dir := filepath.Join(dirRoot, "db", name)
			_, err := os.Stat(dir)

			if os.IsNotExist(err) {
				w.WriteHeader(http.StatusNotFound)
				tpl.ExecuteTemplate(w, "404.html", Obj{})
				return
			}

			var doc Doc
			doc.Name = name

			if binary, err := ioutil.ReadFile(filepath.Join(dir, "content.html")); err == nil {
				doc.Content = string(binary)
			}

			o := Obj{}
			o["Doc"] = doc
			tpl.ExecuteTemplate(w, "del.html", o)
		}

		if r.Method == "POST" {
			name := muxPath(route["name"])

			dir := filepath.Join(dirRoot, "db", name)
			_, err := os.Stat(dir)

			if os.IsNotExist(err) {
				error := fmt.Sprintf("document '%s' does not exist", name)
				muxError(error)
				url := path.Join("/del", name)
				http.Redirect(w, r, url, http.StatusSeeOther)
				return
			}

			os.Remove(filepath.Join(dir, "content.html"))
			os.Remove(dir)

			url := path.Join("/all")
			http.Redirect(w, r, url, http.StatusSeeOther)
		}
	})

	Mux.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			tpl.ExecuteTemplate(w, "new.html", Obj{})
		}

		if r.Method == "POST" {
			r.ParseForm()

			name := muxPath(r.FormValue("name"))
			content := r.FormValue("content")

			dir := filepath.Join(dirRoot, "db", name)
			_, err := os.Stat(dir)

			if !os.IsNotExist(err) {
				error := fmt.Sprintf("document '%s' already exists", name)
				muxError(error)
				url := path.Join("/new")
				http.Redirect(w, r, url, http.StatusSeeOther)
				return
			}

			err = os.MkdirAll(dir, 0770)

			if err != nil {
				url := path.Join("/new")
				http.Redirect(w, r, url, http.StatusSeeOther)
				return
			}

			os.Chmod(dir, 0770)
			if filepath.Dir(dir) != "." {
				os.Chmod(filepath.Dir(dir), 0770)
			}

			{
				f, err := os.Create(filepath.Join(dir, "content.html"))
				defer f.Close()

				if err != nil {
					url := path.Join("/new")
					http.Redirect(w, r, url, http.StatusSeeOther)
					return
				}

				many, err := fmt.Fprint(f, content)
				if err != nil || many != len(content) {
					// error
				} else {
					// success
				}
			}

			url := path.Join("/get", name)
			http.Redirect(w, r, url, http.StatusSeeOther)
		}
	})

	muxNotFound := func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusNotFound)
		tpl.ExecuteTemplate(w, "404.html", Obj{})
	}
	Mux.NotFoundHandler = http.HandlerFunc(muxNotFound)
}

func muxStatic(mux *router.Router, relativePath, root string) {
	mux.PathPrefix(relativePath).Handler(http.StripPrefix(relativePath, http.FileServer(http.Dir(root))))
}

func muxError(s string) {
	fmt.Println(s)
}

func muxPath(s string) string {
	return strings.ReplaceAll(s, "/", "-")
}
