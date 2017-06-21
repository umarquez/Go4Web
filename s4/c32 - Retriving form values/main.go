package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tpl *template.Template

type MyResponse struct {
	Method        string
	URL           *url.URL
	Submissions   map[string][]string
	Header        http.Header
	Host          string
	ContentLength int64
}

type HTTPManager struct{}

// ServeHTTP implements the http.Handler interface.
// It contains a function ServeHTTP with a ResponseWriter an a pointer to an
// http.Request.
func (manager HTTPManager) ServeHTTP(response http.ResponseWriter, req *http.Request) {
	err := req.ParseForm() // Getting Form values (GET and POST)
	if err != nil {
		log.Fatal(err)
	}

	data := MyResponse{
		Method:        req.Method,
		URL:           req.URL,
		Submissions:   req.Form,
		Header:        req.Header,
		Host:          req.Host,
		ContentLength: req.ContentLength,
	}

	// Setting headers to the response
	response.Header().Set("umarquez", "just for me.")

	// Send data to the template
	tpl.ExecuteTemplate(response, "index.gohtml", data)
}

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	var httpManager HTTPManager
	http.ListenAndServe(":8080", httpManager)
}
