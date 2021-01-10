package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	john := person{
		firstName: "John",
		lastName:  "Doe",
		contact: contactInfo{
			email:   "doeJohn@bigCorp.com",
			zipCode: 9400,
		},
	}
	var peter person
	var contact contactInfo

	peter.firstName = "Peter"
	peter.lastName = "Parker"

	contact.email = "PeterParker@spiderman.com"
	contact.zipCode = 1000

	peter.contact = contact

	fmt.Println(john)
	fmt.Printf("%+v", peter) // Prints the keys and values
}
