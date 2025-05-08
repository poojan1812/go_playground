// Struct is kind of a collection for different types.
// Generally declare things that go together buy types are different. i.e - latitude and longitude
package main

import (
	"encoding/json"
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
		Hours   float64 `json:"HoursJson"`
		Minutes float64
	}

	timenow := time{5.0, 30.0}

	// Structures are copied. Changes in one does not apply to other struct. (Similar to arrays)

	// Slice of structure - Create a slice where each value is a structure of key value pair.
	type student struct {
		Name    string
		ID      float64
		Subject string
	}

	pricing := []student{
		{Name: "ABC", ID: 1.0, Subject: "Science"},
		{Name: "DEF", ID: 2.0, Subject: "Mathematics"},
		{Name: "GHI", ID: 3.0, Subject: "Cyber Security"},
	}
	fmt.Println(pricing)

	// JSON is a standard data format and commonly used for storing and sharing data in a structured way.
	// JSON Marshal function encodes the given data, and returns encoded version with the error (if any)

	// JSON module in GO expects the initial uppercase and CamelCase for multi-word variables or types. (Fields must be exported for the JSON package to see them) i.e line 37 type declaration.

	// Struct Tags: Custom JSON output can be set by declaring the respective key for JSON in the struct declaration. i.e line 37 type declaration.
	encoded, error := json.Marshal(timenow)
	fmt.Println(timenow)
	fmt.Println(string(encoded), error)

}
