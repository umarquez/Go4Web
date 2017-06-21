package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func catch(e error) {
	if e != nil {
		log.Fatalln(e)
	}
} // END catch

func init() {
	tpl = template.Must(template.ParseFiles("template.gohtml"))
} // END init

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of  no beliefs",
	}

	err := tpl.Execute(os.Stdout, buddha)
	catch(err)
} // END main
