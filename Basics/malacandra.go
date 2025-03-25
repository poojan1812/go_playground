/*
Malacandra is another name for Mars in The Space Trilogy by C. S. Lewis. Write a program
to determine how fast a ship would need to travel (in km/h) in order to reach Malacandra
in 28 days. Assume a distance of 56,000,000 km.
*/
package main

import (
	"fmt"
)

func main() {
	var distance = 56000000
	var time = 28 * 24

	fmt.Printf("Required speed in km/h is %v\n", distance/time)
}
