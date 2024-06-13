package main

import "fmt"

// interface type
type Animal interface {
	Says() string
	HowManyLegs() int
}

// Dog is the type for dogs
type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

// Says is required by the Animal interface
func (d *Dog) Says() string {
	return d.Sound
}

// HowManyLegs is required by the Animal interface
func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

// Cat is the type for cats
type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

// Says is required by the Animal interface
func (c *Cat) Says() string {
	return c.Sound
}

// HowManyLegs is required by the Animal interface
func (c *Cat) HowManyLegs() int {
	return c.NumberOfLegs
}

func main() {
	// create a variable of type Dog
	dog := Dog{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}

	// Pass dog to riddle. Although dog is of type Dog, it satisifies the
	// interface requirements for Animal because it implements all of Animal's required functions.
	Riddle(&dog)

	// Create  variable of type Cat
	var cat Cat
	cat.Name = "cat"
	cat.NumberOfLegs = 4
	cat.Sound = "meow"
	cat.HasTail = true

	// Pass cat to riddle. Although dog is of type Cat, it satisifies the
	// interface requirements for Animal because it implements all of Animal's required functions.
	Riddle(&cat)
}

// Riddle takes a parameter of type Animal, but will accept both Dog and Cat, since both of those types
// satisfy the interface requirements for Animal, because they both have the correct functions.
func Riddle(a Animal) {
	riddle := fmt.Sprintf(`This animal says "%s" and has %d legs. What animal is it?`, a.Says(), a.HowManyLegs())
	fmt.Println(riddle)
}
