/*
Mars Ticket Generator

This program generates a table of 10 random tickets for space travel to Mars.
Each ticket includes:
  - Spaceline (randomly chosen from Space Adventures, SpaceX, Virgin Galactic)
  - Trip duration in days (based on speed and distance to Mars: 62,100,000 km)
  - Trip type (One-way or Round-trip)
  - Price in millions of dollars ($36M–$50M for one-way, double for round-trip)

The program uses random values for speed (16–30 km/s) and trip type, and outputs
a formatted table of ticket details.
*/

package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	const distance = 62100000 // Distance to Mars in kilometers

	// Split spaceline and trip type strings into slices for random selection
	var spaceline_split = strings.Split("Space Adventures,SpaceX,Virgin Galatic", ",")
	var trip_type_split = strings.Split("One-way,Round-trip", ",")

	// Print table header
	fmt.Printf("%-20s %-10s %-15s %-10s\n", "Spaceline", "Days", "Trip-type", "Price")
	fmt.Println("=====================================================")

	// Generate 10 random tickets
	for random := 0; random < 10; random++ {
		var speed = rand.Intn(15) + 16                // Random speed between 16 and 30 km/s
		var price = 36 + (speed - 16)                 // Base price increases with speed
		var duration = (distance / speed) / 86400     // Calculate trip duration in days
		var spaceline = spaceline_split[rand.Intn(3)] // Randomly select a spaceline
		var trip_type = trip_type_split[rand.Intn(2)] // Randomly select trip type

		// Double the price for round-trip tickets
		if trip_type == "Round-trip" {
			price *= 2
		}

		// Print ticket details in a formatted table row
		fmt.Printf("%-20s %-10d %-15s $%-10d\n", spaceline, duration, trip_type, price)
	}
}
