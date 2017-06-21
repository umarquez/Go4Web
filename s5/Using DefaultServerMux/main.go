package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Something from the dog...")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Mew!")
}

/*
 * If we use the http.HandleFunc it'll append the function to the pattern we
 * set as parameter.
 * We could also use the http.Handle to append an http.Handler instance to the
 * selected pattern to the DefaultServerMux.
 */
func main() {
	// http.Handle("pattern". http.Handler instance)

	// Catches "/dog/*" using the DefaultServerMux
	http.HandleFunc("/dog/", d)
	// Catches "/cat" only using the DefaultServerMux
	http.HandleFunc("/cat", c)

	// Using nil as handler uses the DefaultServerMux
	http.ListenAndServe(":8080", nil)
}
