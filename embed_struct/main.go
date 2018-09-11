package main

import (
	"fmt"
)

type contacInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contacInfo
}

// another way of embeding stuct is
/*
type person struct {
	firstName string
	lastName  string
	contacInfo
}

but calling for this syntax will be
jim := person{
		firstName: "jim",
		lastName:  "Alex",
		contacInfo{
			email:   "jimmu@gym.com",
			zipCode: 260150,
		},
	}

*/

func main() {
	jim := person{
		firstName: "jim",
		lastName:  "Alex",
		contact: contacInfo{
			email:   "jimmu@gym.com",
			zipCode: 260150,
		},
	}
	// jimPointer := &jim
	//above code is address of jim
	// below is go short cut for &jim
	jim.updateName("Thummy")
	jim.print()
}

// signal and reciever in go

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
func (p person) print() {
	fmt.Printf("%+v", p)
}
