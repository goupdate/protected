package main

import (
	"fmt"
	"protected"
)

func main() {
	// Initialize with an integer
	protInt := protected.New(10)

	// Get the value
	val := protInt.Get()
	fmt.Println("Initial value:", val)

	// Set a new value
	protInt.Set(20)
	fmt.Println("Updated value:", protInt.Get())

	// Update the value using a function
	protInt.Update(func(val int) int {
		return val * 2
	})
	fmt.Println("Doubled value:", protInt.Get())

	// Using DoWithLock to safely perform actions on the value
	protInt.DoWithLock(func(val *int) {
		*val = *val + 10
	})
	fmt.Println("Incremented value:", protInt.Get())

	// Using DoWithRLock to safely read the value
	protInt.DoWithRLock(func(val *int) {
		fmt.Println("Read value:", *val)
	})
}
