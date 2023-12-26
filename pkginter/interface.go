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

// Define a struct 'circle' with a 'radius' field
type circle struct {
	radius float64
}

// Implement the 'area' method for the 'circle' struct
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Implement the 'perim' method for the 'circle' struct
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Example function demonstrating the use of the 'geometry' interface
func InterfaceExample1() {
	// Create an instance of the 'circle' struct
	c := circle{radius: 5}

	// Call the 'area' and 'perim' methods using the 'geometry' interface
	// The type 'circle' automatically satisfies the 'geometry' interface
	fmt.Println("Area:", c.area())
	fmt.Println("Perimeter:", c.perim())
}
