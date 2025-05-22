// Composing structures simply means combining multiple small structures and extend their capabilities.

package main

import (
	"fmt"
)

func main() {
	type bank struct {
		name        string
		location    string
		branch_code int
	}
	type address struct {
		street_number string
		city          string
		postal_code   float64
	}
	type branch_details struct {
		transit_number int
		bank           bank
		address        address
	}
	rbc := bank{"rbc", "stc", 0011}
	rbc_address := address{"300", "Toronto", 123.0}

	// Extended `branch_details` type to combine other two structs `bank` and `address`
	branch_detail := branch_details{transit_number: 01, bank: rbc, address: rbc_address}
	fmt.Println(branch_detail)
	// Specific detail can also be accessed
	fmt.Printf("location for the branch: %v\n", branch_detail.bank.location)

	// Method forwarding makes it convenient to use multiple methods.
	// Go Supports this with Struct embedding concept. Specify a type without field name to embed it in a struct.
	type report struct {
		temperature
		sol int
		area
	}
	// Here, all the methods on temperature type are automatically accesible to report.
	// Fields of inner structure is accesible to the outer structure.
	// report.temperature.high can also be accessed through report.high ~ change in either of this is applicable to the other as it shares the same underlying value

	// Go compiler reports `name collision` error if two or more types are embedded with the same method.
	// Here, go does not know which method with which type to forward the call.
	// To solve this, embed the main type with a method and manually forward methods to the embedded types.
	// i.e return r.sol.days(s2) ~ here report type is manually forwarded to days method using sol type.

}
