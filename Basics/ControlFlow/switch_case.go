package main

import (
	"fmt"
)

// Switch compares all cases in value stored in level variable
func main() {
	fmt.Println("Learning golang")
	var level = "Rookie programmer"

	switch level {
	case "Advanced":
		fmt.Println("No need to learn basics")
	case "Rookie programmer", "Don't have experience":
		fmt.Println("Start from the basics and take notes")
		fallthrough // falls through the next case even if it's not a match
	default:
		fmt.Println("And learn something new everyday")
	}
}
