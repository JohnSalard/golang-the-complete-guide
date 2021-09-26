package main

import "fmt"

type contactInfo struct {
	mobile  string
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	age       uint16
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(firstName string) {
	p.firstName = firstName
}

func updateSlice(s []string, value string) {
	s[0] = value
}

func main() {
	por1 := person{
		firstName: "Chumphol",
		lastName:  "Benjamitr",
		age:       31,
		contactInfo: contactInfo{
			mobile:  "0995051168",
			email:   "Chumphol.bjm@gmail.com",
			zipCode: 11000,
		},
	}

	mySlice := []string{"Zero", "Two", "Three"}
	fmt.Println(mySlice)
	updateSlice(mySlice, "One")

	fmt.Println(mySlice)

	por1.updateName("Hello")
	por1.print()
	por1.updateName("Test")
	por1.print()

}
