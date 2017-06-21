package main

import (
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

func main() {
	fs := http.FileServer(http.Dir("starting-files/"))
	http.Handle("/", fs)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	}
}
