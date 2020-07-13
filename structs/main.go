package main

import "fmt"

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
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alexanderson@gmail.com",
			zipCode: 100,
		},
	}

	alex.updateName("Alexa")
	alex.print()
	fmt.Println(*&alex)

}

func (pointerToPerson *person) updateName(newFisrtName string) {
	(*pointerToPerson).firstName = newFisrtName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
