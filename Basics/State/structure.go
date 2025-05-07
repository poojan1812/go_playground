// Struct is kind of a collection for different types.
// Generally declare things that go together buy types are different. i.e - latitude and longitude
package main

import (
	"fmt"
)

func main() {
	var shopping struct {
		name  string
		price int
	}

	shopping.name = "Apple"
	shopping.price = 10
	fmt.Println(shopping.name, shopping.price)

	// If multiple structures are needed with the same type, consider declaring custom type struct
	type location struct {
		latitude, longitude float64
	}

	// The custom type declared above is used twice below.
	var first location
	first.latitude = 1.0
	first.longitude = 1.1
	fmt.Println(first.longitude, first.latitude)

	var second location
	second.latitude = 2.0
	second.longitude = 2.1
	fmt.Println(second.latitude, second.longitude)

	// COMPOSITE LITERAL DECLARATION
	type time struct {
		hours, minutes float64
	}

	timenow := time{hours: 5.0, minutes: 30.0}
	fmt.Println(timenow)

	// Structures are copied. Changes in one does not apply to other struct. (Similar to arrays)

	// Slice of structure -

}
