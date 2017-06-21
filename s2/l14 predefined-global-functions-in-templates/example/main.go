package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type user struct {
	Name   string
	Email  string
	Admin  bool
	Active bool
	Access []string
}

func catch(e error) {
	if e != nil {
		log.Fatalln(e)
	}
} // END catch

func init() {
	tplName := "template.gohtml"
	tpl = template.Must(template.ParseFiles(tplName))
} // END init

func main() {
	users := []user{}

	users = append(users, user{
		Name:   "root",
		Email:  "root@system.local",
		Admin:  true,
		Active: true,
		Access: []string{"add", "remove", "modify", "read"},
	})

	users = append(users, user{
		Name:   "supervisor",
		Email:  "super@system.local",
		Admin:  true,
		Active: true,
		Access: []string{"read", "modify"},
	})

	users = append(users, user{
		Name:   "user",
		Email:  "user@system.local",
		Admin:  false,
		Active: true,
		Access: []string{},
	})

	users = append(users, user{
		Name:   "reporter",
		Email:  "reports@system.local",
		Admin:  false,
		Active: true,
		Access: []string{},
	})

	users = append(users, user{
		Name:   "test",
		Email:  "",
		Admin:  false,
		Active: true,
		Access: []string{},
	})

	users = append(users, user{
		Name:   "inactive",
		Email:  "null@system.local",
		Admin:  false,
		Active: false,
		Access: []string{},
	})

	err := tpl.Execute(os.Stdout, users)
	catch(err)
} // END main
