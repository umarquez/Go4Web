package main

import (
	"io"
	"net/http"
)

type hotdog struct{}

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Something from the dog...")
}

type hotcat struct{}

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Mew!")
}

func main() {
	// HTTP Mux
	mux := http.NewServeMux()
	// Catches "/dog/*"
	mux.Handle("/dog/", hotdog{})
	// Catches "/cat" only
	mux.Handle("/cat", hotcat{})

	http.ListenAndServe(":8080", mux)
}
