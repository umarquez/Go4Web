package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func catch(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func init() {
	tpl = template.Must(template.ParseFiles("template.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", `Release self-focus; embrace other-focus`)
	catch(err)
}
