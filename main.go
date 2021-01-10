package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	john := person{firstName: "John", lastName: "Doe"}
	var peter person
	peter.firstName = "Peter"
	peter.lastName = "Parker"
	fmt.Println(john)
	fmt.Printf("%+v", peter) // Prints the keys and values
}
