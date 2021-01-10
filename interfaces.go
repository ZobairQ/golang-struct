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

func (dog) speak() {
	fmt.Println("Woof!")
}

func (cat) speak() {
	fmt.Println("Meow")
}

func letAnimalSpeak(a animal) {
	a.speak()
}
