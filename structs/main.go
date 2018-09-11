package main

import "fmt"

// tells go about new custom type with help of struct

type person struct {
	firstName string
	lastName  string
}

func main() {
	//  first way to create struct is alex := person{"Alex","Anderson"} but this
	//  this method is wierd best method is as follows
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// another method for declaring a struct as empty
	// var alex person
	// another way
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
