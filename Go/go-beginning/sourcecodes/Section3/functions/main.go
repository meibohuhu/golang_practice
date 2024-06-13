package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)

// interface type

// Animal is a type, with three fields
type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

// Says is a receiver function, tied to the type Animal
// Because it is tied to the type, you can only call it
// when using a variable of type Animal
func (a *Animal) Says() {
	fmt.Printf("A %s says %s", a.Name, a.Sound)
	fmt.Println()
}

// HowManyLegs is a function tied to the Animal type
func (a *Animal) HowManyLegs() {
	fmt.Printf("A %s has %d legs", a.Name, a.NumberOfLegs)
	fmt.Println()
}

func main() {
	var dog Animal
	dog.Name = "dog"
	dog.Sound = "woof"
	dog.NumberOfLegs = 4
	dog.Says()

	cat := Animal{
		Name:         "cat",
		Sound:        "meow",
		NumberOfLegs: 4,
	}

	cat.Says()
	cat.HowManyLegs()
}

// addTwoNumbers returns the sum of 2 ints
func addTwoNumbers(x, y int) int {
	return x + y
}

// naked return - rarely used, but possible, and only
// intended for short functions.
func add(x, y int) (sum int) {
	sum = x + y
	return
}

// variadic function - takes variable number of parameters
func sumMany(nums ...int) int {
	total := 0
	for _, x := range nums {
		total = total + x
	}

	return total
}
