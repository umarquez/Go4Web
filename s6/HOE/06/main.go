package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

/*
 * Starting with the code in the "starting-files" folder, wire this program up
 * so that it works:
 * - ParseGlob in an init function.
 * - Use HandleFunc for each of the routes.
 * - Combine apply & applyProcess into one func called "apply".
 * - Inside the func "apply", use this code to create the logic to respond
 *   differently to a POST method and a GET method.
 *
 *     if req.Method == http.MethodPost {
 *       // code here
 *       return
 *     }
 */

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	fmt.Printf("%#v", tpl)

	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/apply", apply)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func apply(w http.ResponseWriter, req *http.Request) {
	var err error
	if req.Method == http.MethodPost {
		err = tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
		HandleError(w, err)
	} else {
		err = tpl.ExecuteTemplate(w, "apply.gohtml", nil)
		HandleError(w, err)
	}

	if err != nil {
		log.Panic(err)
	}

	return
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
