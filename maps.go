package main

import "fmt"

func newMap() {
	var words map[string]string // declaration
	words = map[string]string{} // new map()

	wordsCreatedWithMake := make(map[string]string) // new map

	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#0000FF",
		"green": "#00FF00",
	}

	words["go"] = "A programming language by Google"
	wordsCreatedWithMake["go"] = "A programming language by Google"

	delete(words, "go") // deletes by key

	// fmt.Println(colors)
	// fmt.Println(words)
	// fmt.Println(wordsCreatedWithMake)
	printMap(colors)
}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}
