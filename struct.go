package main

import (
	"fmt"
)

type user struct {
	name    string
	email   string
	phone   int
	address string
}

func structExample1() {
	//first Way to initialize value in Struct
	u := user{"tinku", "xyz@gmail.com", 12, "xyz"}
	fmt.Println(u)

	//Second Way to initialize value in Struct
	u1 := user{name: "Tanish ", email: "tinku1234@gmail.com", phone: 12, address: "asdfghj"}
	//partial and other will come default
	u3 := user{name: "Tanish ", email: "tinku1234@gmail.com"}

	//access value by dot operator
	fmt.Println(u1.name, u1.address, u1.phone, u1.email, u3.name)

	//structs with pointer
	u4 := &user{name: "Tanish Bhnadari", email: "tinku@gmail.com"}
	//Both way supported to access variables
	fmt.Println((*u4).name)
	fmt.Println(u4.name, u3)
}

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func structExample2() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
