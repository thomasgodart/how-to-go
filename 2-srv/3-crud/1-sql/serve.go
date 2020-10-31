package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
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

			var dbDocs []DBDoc
			db.Find(&dbDocs)

			docs := make(map[string]Doc)

			for _, dbDoc := range dbDocs {

				docs[dbDoc.Name] = Doc{
					URL:  path.Join("/", "get", muxPath(dbDoc.Name)),
					Name: dbDoc.Name,
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

			var dbDoc DBDoc
			db.First(&dbDoc, "name = ?", name)

			if dbDoc.ID == 0 {
				w.WriteHeader(http.StatusNotFound)
				tpl.ExecuteTemplate(w, "404.html", Obj{})
				return
			}

			var doc Doc
			doc.Name = dbDoc.Name
			doc.Content = dbDoc.Content

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

			var dbDoc DBDoc
			db.First(&dbDoc, "name = ?", name)

			if dbDoc.ID == 0 {
				w.WriteHeader(http.StatusNotFound)
				tpl.ExecuteTemplate(w, "404.html", Obj{})
				return
			}

			var doc Doc
			doc.Name = dbDoc.Name
			doc.Content = dbDoc.Content

			o := Obj{}
			o["Doc"] = doc
			tpl.ExecuteTemplate(w, "set.html", o)
		}

		if r.Method == "POST" {
			r.ParseForm()

			name := muxPath(r.FormValue("name"))
			content := r.FormValue("content")

			var dbDoc DBDoc
			db.First(&dbDoc, "name = ?", name)

			if dbDoc.ID == 0 {
				error := fmt.Sprintf("document '%s' does not exist", name)
				muxError(error)
				url := path.Join("/set", name)
				http.Redirect(w, r, url, http.StatusSeeOther)
				return
			}

			db.Model(&dbDoc).Update("content", content)

			url := path.Join("/get", name)
			http.Redirect(w, r, url, http.StatusSeeOther)
		}
	})

	Mux.HandleFunc("/del/{name}", func(w http.ResponseWriter, r *http.Request) {
		route := router.Vars(r)

		if r.Method == "GET" {
			name := muxPath(route["name"])

			var dbDoc DBDoc
			db.First(&dbDoc, "name = ?", name)

			if dbDoc.ID == 0 {
				w.WriteHeader(http.StatusNotFound)
				tpl.ExecuteTemplate(w, "404.html", Obj{})
				return
			}

			var doc Doc
			doc.Name = name
			doc.Content = dbDoc.Content

			o := Obj{}
			o["Doc"] = doc
			tpl.ExecuteTemplate(w, "del.html", o)
		}

		if r.Method == "POST" {
			name := muxPath(route["name"])

			var dbDoc DBDoc
			db.First(&dbDoc, "name = ?", name)

			if dbDoc.ID == 0 {
				error := fmt.Sprintf("document '%s' does not exist", name)
				muxError(error)
				url := path.Join("/del", name)
				http.Redirect(w, r, url, http.StatusSeeOther)
				return
			}

			db.Delete(&dbDoc, dbDoc.ID)

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

			var dbDoc DBDoc
			db.First(&dbDoc, "name = ?", name)

			if dbDoc.ID != 0 {
				error := fmt.Sprintf("document '%s' already exists", name)
				muxError(error)
				url := path.Join("/new")
				http.Redirect(w, r, url, http.StatusSeeOther)
				return
			}

			db.Create(&DBDoc{Name: name, Content: content})

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
