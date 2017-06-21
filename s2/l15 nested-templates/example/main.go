package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func catch(e error) {
	if e != nil {
		log.Fatalln(e)
	}
} // END catch

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
} // END init

func main() {
	data := "<p>\"Hello World!\"</p>"

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", data)
	catch(err)
} // END main
