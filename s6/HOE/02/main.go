package main

import (
	"log"
	"net/http"
)

/*
 * Serve the files in the "starting-files" folder, use "http.FileServer"
 */

func main() {
	http.Handle("/", http.FileServer(http.Dir("./starting-files")))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	}
}
