package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	router "github.com/gorilla/mux"
)

type Obj map[string]interface{}

func init() {

	Mux.HandleFunc("/all", func(w http.ResponseWriter, r *http.Request) {

		{

			var dbDocs []DBDoc
			db.Find(&dbDocs)

			o := Obj{}
			o["all"] = dbDocs
			o["len"] = len(dbDocs)
			b, _ := json.Marshal(o)
			w.Write(b)
		}
	})

	Mux.HandleFunc("/get/{name}", func(w http.ResponseWriter, r *http.Request) {
		route := router.Vars(r)

		{
			name := muxPath(route["name"])

			var dbDoc DBDoc
			db.First(&dbDoc, "name = ?", name)

			if dbDoc.ID == 0 {
				w.WriteHeader(http.StatusNotFound)
				error := fmt.Sprintf("document '%s' does not exist", name)

				o := Obj{}
				o["error"] = error
				b, _ := json.Marshal(o)
				w.Write(b)

				return
			}

			o := Obj{}
			o["get"] = dbDoc
			b, _ := json.Marshal(o)
			w.Write(b)
		}
	})

	Mux.HandleFunc("/set/{name}", func(w http.ResponseWriter, r *http.Request) {
		route := router.Vars(r)

		{
			r.ParseForm()

			name := muxPath(route["name"])
			content := r.FormValue("content")

			var dbDoc DBDoc
			db.First(&dbDoc, "name = ?", name)

			if dbDoc.ID == 0 {
				w.WriteHeader(http.StatusNotFound)
				error := fmt.Sprintf("document '%s' does not exist", name)

				o := Obj{}
				o["error"] = error
				b, _ := json.Marshal(o)
				w.Write(b)

				return
			}

			db.Model(&dbDoc).Update("content", content)

			db.First(&dbDoc, "name = ?", name)

			o := Obj{}
			o["success"] = true
			o["set"] = dbDoc
			b, _ := json.Marshal(o)
			w.Write(b)
		}
	})

	Mux.HandleFunc("/del/{name}", func(w http.ResponseWriter, r *http.Request) {
		route := router.Vars(r)

		{
			name := muxPath(route["name"])

			var dbDoc DBDoc
			db.First(&dbDoc, "name = ?", name)

			if dbDoc.ID == 0 {
				w.WriteHeader(http.StatusNotFound)
				error := fmt.Sprintf("document '%s' does not exist", name)

				o := Obj{}
				o["error"] = error
				b, _ := json.Marshal(o)
				w.Write(b)

				return
			}

			db.Delete(&dbDoc, dbDoc.ID)

			o := Obj{}
			o["success"] = true
			b, _ := json.Marshal(o)
			w.Write(b)
		}
	})

	Mux.HandleFunc("/new/{name}", func(w http.ResponseWriter, r *http.Request) {
		route := router.Vars(r)

		{
			r.ParseForm()

			name := muxPath(route["name"])
			content := r.FormValue("content")

			var dbDoc DBDoc
			db.First(&dbDoc, "name = ?", name)

			if dbDoc.ID != 0 {
				error := fmt.Sprintf("document '%s' already exists", name)

				o := Obj{}
				o["error"] = error
				b, _ := json.Marshal(o)
				w.Write(b)

				return
			}

			db.Create(&DBDoc{Name: name, Content: content})

			db.First(&dbDoc, "name = ?", name)

			o := Obj{}
			o["success"] = true
			o["new"] = dbDoc
			b, _ := json.Marshal(o)
			w.Write(b)
		}
	})

	muxNotFound := func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusBadRequest)
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
