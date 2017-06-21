package main

import (
	"html/template"
	"log"
	"net/http"
)

/*
 * Serve the files in the "starting-files" folder
 * To get your images to serve, use only this:
 *   fs := http.FileServer(http.Dir("public"))
 *
 * Hint: look to see what type FileServer returns, then think it through.
 */

var tpl *template.Template

func index(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}

func init() {
	tpl = template.Must(template.ParseFiles("starting-files/templates/index.gohtml"))
}

func main() {
	fs := http.FileServer(http.Dir("starting-files/public"))
	http.HandleFunc("/", index)
	http.Handle("/pics/", fs)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	}
}
