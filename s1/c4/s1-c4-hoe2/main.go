package main

import "fmt"

/*
 * HANDS ON 2
 * create a struct that holds person fields
 * create a struct that holds secret agent fields and embeds person type
 * attach a method to person: pSpeak
 * attach a method to secret agent: saSpeak
 * create a variable of type person
 * create a variable of type secret agent
**/

type person struct {
	name string
}

func (p person) pSpeak() {
	fmt.Println("Hello, my name is", p.name)
}

type secretAgent struct {
	person
	code string
}

func (sa secretAgent) saSpeak() {
	fmt.Println("I'm the agent", sa.code, " Name:", sa.person.name)
}

func main() {
	me := person{"@umarquez"}
	jb := secretAgent{
		person{"James Bond"},
		"007",
	}

	me.pSpeak()
	jb.saSpeak()
}
