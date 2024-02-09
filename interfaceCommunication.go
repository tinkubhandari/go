// Go program to illustrate the concept
// of the embedding interfaces
package main

import "fmt"

// Interface 1
type manWalk interface {
	walk()
}

// Interface 2
type manSpeak interface {
	speak()
}

// Interface 3

// Interface 3 embedded with
// interface 1 and 2
type manFunctionality interface {
	manWalk
	manSpeak
	//walk()
	//speak()
}

type man struct {
	name string
	age  int
}

// Implementing method of
// the interface 1
func (m man) speak() {
	fmt.Println("Tinku is Speaking")
}

// Implementing method of
// the interface 2
func (m man) walk() {
	fmt.Println("Tinku is Walking")
}

func InterfaceCommunication2() {
	// Assigning values
	// to the structure
	person12 := man{"tinku", 30}

	// Accessing the methods of
	// the interface 1 and 2
	// Using manFunctionality interface
	var f manFunctionality = person12
	f.speak()
	f.walk()
}

// As shown in the above example we have three interfaces. Interface 1 and 2 are simple interfaces and interface 3 is the embedded interface which holds both 1 and 2 interfaces in it. So if any changes take place in the interface 1 and interface 2 will reflect in interface 3. And interface 3 can access all the methods present in interface 1 and 2. Example 2:

// Interface 1
type AuthorDetails interface {
	details()
}

// Interface 2
type AuthorArticles interface {
	articles()
}

// Interface 3

// Interface 3 embedded with
// interface 1 and 2
type FinalDetails interface {
	AuthorDetails
	AuthorArticles
}

// Structure
type author struct {
	a_name    string
	branch    string
	college   string
	year      int
	salary    int
	particles int
	tarticles int
}

// Implementing method of
// the interface 1
func (a author) details() {

	fmt.Printf("Author Name: %s", a.a_name)
	fmt.Printf("\nBranch: %s and passing year: %d",
		a.branch, a.year)
	fmt.Printf("\nCollege Name: %s", a.college)
	fmt.Printf("\nSalary: %d", a.salary)
	fmt.Printf("\nPublished articles: %d", a.particles)
}

// Implementing method
// of the interface 2
func (a author) articles() {

	pendingarticles := a.tarticles - a.particles
	fmt.Printf("\nPending articles: %d", pendingarticles)
}

// Main value
func InterfaceCommunication() {

	// Assigning values
	// to the structure
	values := author{
		a_name:    "Mickey",
		branch:    "Computer science",
		college:   "XYZ",
		year:      2012,
		salary:    50000,
		particles: 209,
		tarticles: 309,
	}

	// Accessing the methods of
	// the interface 1 and 2
	// Using FinalDetails interface
	var f FinalDetails = values
	f.details()
	f.articles()
}


//3RD CASE ---------------------------------------------------------->
// Explanation: As shown in the above example we have three interfaces. 
//Interface 1 and 2 are simple interfaces and interface 3 is the embedded interface which holds both 1 and 2 interfaces method only signatures in it. 
//So if any changes take place in the interface 1 and interface 2’s methods it will not reflect in interface 3 until we define explicitly in interface 3. 
//And interface 3 can access all the methods present in interface 1 and 2 only if those mentioned in interface 3. Example 3:

// type FinalDetails interface {
// 	details()
// 	articles()
// }


//4TH CASE ----------------------------------------------------------->
// Explanation: As shown in the above example we have three interfaces. 
//Interface 1 and 2 are simple interfaces and interface 3 is the embedded interface which holds both interface 1’s method signatures, 
//interface 2 and it’s own method in it. So if any changes take place in the interface 2’s method it will reflect in interface 3. 
//And interface 3 can access all the methods present in it including interface 2.
//We can only access interface 1’s method which signature define in interface3.

// type FinalDetails interface {
//     details()---> this changes will not reflect in interface 1
//     AuthorArticles ----> this chnages will reflect
//     cdeatils()---> this changes will not reflect in interface 2
// }
