package main

import "fmt"

func newMap() {
	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#0000FF",
		"green": "#00FF00",
	}
	
	fmt.Println(colors)
}
