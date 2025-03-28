package main

import (
	"fmt"
)

// if else loop demo
func main() {
	var direction = "Run east"

	if direction == "Run east" {
		fmt.Println("Keep running till you find me")
	} else if direction == "Run west" {
		fmt.Println("Stop now")
	} else {
		fmt.Println("Do nothing")
	}
}
