package main

import (
	"html/template"
	"log"
	"net/http"
)

/*
 * Serve the files in the "starting-files" folder
 * To get your images to serve, use:
 *   func StripPrefix(prefix string, h Handler) Handler
 *   func FileServer(root FileSystem) Handler
 *
 * Constraint: you are not allowed to change the route being used for images in
 * the template file.
 */

var tpl *template.Template

func index(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}

func init() {
	tpl = template.Must(template.ParseFiles("starting-files/templates/index.gohtml"))
}

func main() {
	fs := http.FileServer(http.Dir("starting-files/public/"))
	http.HandleFunc("/", index)
	http.Handle("/resources/", http.StripPrefix("/resources", fs))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	}
}
