package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

// create a FuncMap to register functions.
// "uc" is what the func will be called in the template
// "uc" is the ToUpper func from package strings
// "ft" is a func I declared
// "ft" slices a string, returning the first three characters
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func catch(e error) {
	if e != nil {
		log.Fatalln(e)
	}
} // END catch

func init() {
	tplName := "template.gohtml"

	// First, we need to create a *template with template.New()
	// then, we neet to assign the FuncMap object to our *templ|ate
	// and before that, we parse the file with ParseFile function.
	// This is beacuse we need to have the function declared before parse the file
	// because we can't use something if it's not declared.
	tpl = template.Must(template.New(tplName).Funcs(fm).ParseFiles(tplName))
} // END init

func firstThree(s string) string {
	// cleanning first and last spaces
	s = strings.TrimSpace(s)
	// getting the first three elements
	s = s[:3]

	return s
}

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
