package main

import (
	"log"
	"os"
	"text/template"
)

func catch(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func main() {
	tpl := template.Must(template.ParseGlob("templates/*"))

	err = tpl.Execute(os.Stdout, nil)
	catch(err)

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	catch(err)

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	catch(err)

	err = tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	catch(err)
}
