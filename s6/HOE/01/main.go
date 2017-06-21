package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

/*
 * - ListenAndServe on port 8080 of localhost.
 * - For the default route "/" Have a func called "foo" which writes to the
 *   response "foo ran".
 * - For the route "/dog/" Have a func called "dog" which parses a template
 *   called "dog.gohtml" and writes to the response "This is from dog" and also
 *   shows a picture of a dog when the template is executed.
 * - Use "http.ServeFile" to serve the file "dog.jpeg"
 */

var tpl *template.Template

func foo(res http.ResponseWriter, req *http.Request) {
	fmt.Println("- Serving foo.")
	content := "<h1>Foo ran</h1>"
	io.WriteString(res, content)
}

func dog(res http.ResponseWriter, req *http.Request) {
	fmt.Println("- Serving dog.")
	content := "This is from dog"
	tpl.ExecuteTemplate(res, "dog.gohtml", content)
}

func dogPic(res http.ResponseWriter, req *http.Request) {
	fmt.Println("- Serving assets.")
	http.ServeFile(res, req, "assets/dog.jpeg")
}

func init() {
	tpl = template.Must(template.ParseFiles("templates/dog.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/assets/dog.jpeg", dogPic)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	}
}
