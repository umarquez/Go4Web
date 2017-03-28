package main

import (
	"fmt"
	"math"
)

/*
 * HANDS ON 1
 * create a type square
 * c+reate a type circle
 * attach a method to each that calculates area and returns it
 * create a type shape which defines an interface as anything which has the area method
 * create a func info which takes type shape and then prints the area
 * create a value of type square
 * create a value of type circle
 * use func info to print the area of square
 * use func info to print the area of circle
 **/

type square struct {
	size float64
}

func (s square) area() float64 {
	return math.Pow(s.size, 2)
}

type circle struct {
	radio float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radio, 2)
}

type shape interface {
	area() float64
}

func info(s shape) {
	a := s.area()
	fmt.Println(a)
}

func main() {
	size := 3.0
	s := square{size}
	c := circle{size}

	info(s)
	info(c)
}
