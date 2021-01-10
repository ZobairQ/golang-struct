package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName   string
	lastName    string
	contactInfo // is the same as contactInfo contactInfo
}

func main() {
	dog := dog{}
	letAnimalSpeak(dog)
	newMap()
	john := person{
		firstName: "John",
		lastName:  "Doe",
		contactInfo: contactInfo{
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

	peter.contactInfo = contact

	// Using a address operator
	pointerToPeter := &peter
	pointerToPeter.updateName("Nedri")
	peter.greet()

	// Not using a address operator
	john.updateName("Jane") // go will translate this into a address
	john.greet()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p) // Prints the keys and values
}

func (p *person) greet() {
	fmt.Println("Hello my name is ", p.firstName)
}
