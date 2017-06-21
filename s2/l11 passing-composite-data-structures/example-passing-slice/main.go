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
	tpl = template.Must(template.ParseFiles("template.gohtml"))
} // END init

func main() {
	sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}

	err := tpl.Execute(os.Stdout, sages)
	catch(err)
} // END main
