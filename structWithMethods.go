package main

import (
	"fmt"
)

type Person struct {
	name   string
	rollno int
	age    int
}

// p Person-->RECIEVER
// Methis name--->print
func (p Person) print() {
	fmt.Println(p)
}
func structWithMethod0() {
	p := Person{name: "tinku", rollno: 23, age: 22}
	p.print()
}

// Define a struct named 'rect' with 'width' and 'height' fields
type rect struct {
	width, height int
}

// Method for 'rect' struct: calculates the area
func (r *rect) area() int {
	return r.width * r.height
}

// Method for 'rect' struct: calculates the perimeter
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func structWithMethod1() {
	// Create an instance of the 'rect' struct
	r := rect{width: 10, height: 5}

	// Call the 'area' and 'perim' methods on the 'rect' instance
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Create a pointer to the 'rect' instance
	rp := &r

	// Call the 'area' and 'perim' methods using the pointer
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
