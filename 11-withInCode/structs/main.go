package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	fisrtName string
	lastName  string
	contact   contactInfo
}

func main() {
	//it will not work if firstname and lastname interchange
	//alex := person{"Alex","Anderson"}

	//alternative
	//alex := person{fisrtName: "Alex", lastName: "Anderson"}

	// var alex person
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		fisrtName: "rujal",
		lastName:  "shrestha",
		contact: contactInfo{
			email:   "rujal@gmail.com",
			zipCode: 9744,
		},
	}

	jim.updateName("jimmy")
	jim.print()

}
func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).fisrtName = newFirstName
}
