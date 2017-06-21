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

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	mlk := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases eith heatred but with lovw alone is healed",
	}

	muhammad := sage{
		Name:  "Muhammad",
		Motto: "To overcome evil with good is good, to resist evil by evil is evil",
	}

	sages := []sage{buddha, gandhi, mlk, muhammad}

	err := tpl.Execute(os.Stdout, sages)
	catch(err)
} // END main
