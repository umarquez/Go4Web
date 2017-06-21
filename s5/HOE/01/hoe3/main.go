package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/*
 * 1. Take the previous program in the previous folder and change it so that:
 * - a template is parsed and served
 * - you pass data into the template
 */

var tpl *template.Template

func me(res http.ResponseWriter, req *http.Request) {
	d := []string{"   ____",
		"  / __ \\ _   _ _ __ ___   __ _ _ __ __ _ _   _  ___ ____",
		" / / _` | | | | '_ ` _ \\ / _` | '__/ _` | | | |/ _ \\_  /",
		"| | (_| | |_| | | | | | | (_| | | | (_| | |_| |  __// /",
		" \\ \\__,_|\\__,_|_| |_| |_|\\__,_|_|  \\__, |\\__,_|\\___/___|",
		"  \\____/                              | |",
		"                                      |_|",
	}

	tpl.ExecuteTemplate(res, "template.gohtml", d)
}

func dog(res http.ResponseWriter, req *http.Request) {
	d := []string{
		"^..^      /",
		"/_/\\_____/",
		"   /\\   /\\",
		"  /  \\ /  \\",
	}

	tpl.ExecuteTemplate(res, "template.gohtml", d)
}

func index(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "template.gohtml", []string{"Hola mundo!"})
	if err != nil {
		panic(err)
	}
}

func init() {
	tpl = template.Must(template.ParseFiles("./template.gohtml"))
}

func main() {
	fmt.Println("...")
	http.HandleFunc("/", http.HandlerFunc(index))
	http.HandleFunc("/dog/", http.HandlerFunc(dog))
	http.HandleFunc("/me/", http.HandlerFunc(me))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
