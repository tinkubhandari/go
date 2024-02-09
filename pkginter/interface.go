package pkginter

import (
	"fmt"
	"math"
)

// Define an interface named 'geometry' with methods 'area' and 'perim'
type geometry interface {
	area() float64
	perim() float64
}

// Define a struct 'circle' and 'rectnagle' with a 'radius' field
type rect struct {
    width, height float64
}

type circle struct {
    radius float64
}

func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

// Implement the 'area' method for the 'circle' struct
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Implement the 'perim' method for the 'circle' struct
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

// Example function demonstrating the use of the 'geometry' interface
func InterfaceExample1() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

     // Call the 'area' and 'perim' methods using the 'geometry' interface
     // The type 'circle' automatically satisfies the 'geometry' interface
    measure(r)
    measure(c)
}
