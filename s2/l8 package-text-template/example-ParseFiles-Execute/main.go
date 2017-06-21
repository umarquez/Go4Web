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
	tpl, err := template.ParseFiles("tpl.gohtml")
	catch(err)

	nf, err := os.Create("index.html")
	catch(err)
	defer nf.Close()

	// err = tpl.Execute(os.Stdout, nil) // to the standard out
	err = tpl.Execute(nf, nil) // to a fsile
	catch(err)

	// == Parsing more then one template with tpl as a container
	tpl, err = tpl.ParseFiles("two.gmao", "three.gmao")
	catch(err)

	// == Executing templates
	err = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", nil)
	catch(err)

	err = tpl.ExecuteTemplate(os.Stdout, "two.gmao", nil)
	catch(err)

	err = tpl.ExecuteTemplate(os.Stdout, "three.gmao", nil)
	catch(err)
}
