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

}
