package main

import "fmt"

type Animal interface {
	Says() string
	HowManyLegs() int
	Tail() bool
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

// func (d *Dog) HowManyLegs
func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

func (d *Dog) Tail() bool {
	return true
}

type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

func (c *Cat) Says() string {
	return c.Sound
}
func (c *Cat) HowManyLegs() int {
	return c.NumberOfLegs
}
func (c *Cat) Tail() bool {
	return c.HasTail
}

func main() {
	// create a variable of type Dog
	dog := Dog{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}
	cat := Cat{
		Name:         "cat",
		Sound:        "miao",
		NumberOfLegs: 4,
		HasTail:      false,
	}

	Riddle(&dog) // Dog must implement all functions in interface
	Riddle(&cat)
}

func Riddle(a Animal) {
	riddle := fmt.Sprintf(`This animal says "%s" and has %d legs. has tail %t`, a.Says(), a.HowManyLegs(), a.Tail())
	fmt.Println(riddle)
}
