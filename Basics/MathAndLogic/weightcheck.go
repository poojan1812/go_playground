package main

import (
	"fmt"
)
// printf, println, print functions to print the given text in the line
// printf will provide substituting values form the argument
// printf also provides formatting features for outputs
func main(){
	fmt.Printf("My weight on the surface of mars is %v, and my age would be %v", 200*0.37, 24*365/687)
	fmt.Printf("\nTrying to add padding %25v\n", "it works")
}