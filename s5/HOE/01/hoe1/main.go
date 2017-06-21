package main

import (
	"fmt"
	"io"
	"net/http"
)

/*
 * ListenAndServe on port ":8080" using the default ServeMux.
 * Use HandleFunc to add the following routes to the default ServeMux:
 * "/" "/dog/" "/me/
 * Add a func for each of the routes.
 * Have the "/me/" route print out your name.
 */

func me(res http.ResponseWriter, req *http.Request) {
	d := fmt.Sprintln("   ____")
	d += fmt.Sprintln("  / __ \\ _   _ _ __ ___   __ _ _ __ __ _ _   _  ___ ____")
	d += fmt.Sprintln(" / / _` | | | | '_ ` _ \\ / _` | '__/ _` | | | |/ _ \\_  /")
	d += fmt.Sprintln("| | (_| | |_| | | | | | | (_| | | | (_| | |_| |  __// /")
	d += fmt.Sprintln(" \\ \\__,_|\\__,_|_| |_| |_|\\__,_|_|  \\__, |\\__,_|\\___/___|")
	d += fmt.Sprintln("  \\____/                              | |")
	d += fmt.Sprintln("                                      |_|")
	d += fmt.Sprintln("")

	io.WriteString(res, d)
}

func dog(res http.ResponseWriter, req *http.Request) {
	d := fmt.Sprintln("^..^      /")
	d += fmt.Sprintln("/_/\\_____/")
	d += fmt.Sprintln("   /\\   /\\")
	d += fmt.Sprintln("  /  \\ /  \\")

	io.WriteString(res, d)
}

func index(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hola mundo!")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
