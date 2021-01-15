package main

import "fmt"

/*
The animal interface is bound to dog and cat struct by the receiver functions speak()
*/
type animal interface {
	speak()
}
type dog struct{}
type cat struct{}
type cow struct{}

func (dog) speak() {
	fmt.Println("Woof!")
}

func (cow) speak() {
	fmt.Println("Mooo!")
}

func (cat) speak() {
	fmt.Println("Meow")
}

func letAnimalSpeak(a animal) {
	a.speak()
}
