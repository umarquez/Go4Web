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
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	catch(err)

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	catch(err)

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	catch(err)

	err = tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	catch(err)
}
