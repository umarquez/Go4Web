package main

import "fmt"

/*
 * HANDS ON 3
 * create an interface type that both person and secretAgent implement
 * declare a func with a parameter of the interface’s type
 * call that func in main and pass in a value of type person
 * call that func in main and pass in a value of type secretAgent
**/

type person struct {
	name string
}

func (p person) iAmA() {
	fmt.Printf("I'm a %T\n", p)
}

type secretAgent struct {
	person
	code string
}

func (sa secretAgent) iAmA() {
	fmt.Printf("I'm a %T\n", sa)
}

type something interface {
	iAmA()
}

func showMe(s something) {
	s.iAmA()
}

func main() {
	me := person{"@umarquez"}

	jb := secretAgent{
		person{"James Bond"},
		"007",
	}

	showMe(me)
	showMe(jb)
}
