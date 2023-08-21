package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmial.com",
			zipCode: 94000,
		},
	}

	// No need to make the actual pointer as the receiver can figure it out
	// jimPointer := &jim
	// jimPointer.updateName("jimmy")

	jim.updateName("jimmy")
	jim.print()
	// fmt.Println(jim)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
