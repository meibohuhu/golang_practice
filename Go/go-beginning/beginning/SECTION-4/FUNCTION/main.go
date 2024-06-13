package main

import "fmt"

// Animal is a type, with three fields
type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

// Says is a receiver function, tied to the type Animal
// Because it is tied to the type, you can only call it
// when using a variable of type Animal
func (a *Animal) Says() { //* means we pass in memory of variable, when change NumberOfLegs in the function, the value outside function will be changed
	fmt.Printf("A %s says %s", a.Name, a.Sound)
	fmt.Println()
	a.NumberOfLegs = 7
}

func main() {
	var dog Animal
	dog.Name = "dollar"
	dog.Sound = "wangwang"
	dog.NumberOfLegs = 4
	// println(dog.NumberOfLegs)

	// dog.Says()
	// println(dog.NumberOfLegs)

	cat := Animal{
		Name:         "hhh",
		Sound:        "miaomiao",
		NumberOfLegs: 4,
	}
	cat.Says()
	println(cat.NumberOfLegs)

	a := 10
	modifyByReference(&a)
	println(a)
}

// modify value by reference, pass in pointer/reference to a, deference a and change value to 12
func modifyByReference(a *int) {
	*a = 12
}

func add(x, y int) int {
	sum := x + y
	return sum
}

// variadic function - takes variable number of parameters
func sumMany(nums ...int) int {
	total := 0
	for _, x := range nums {
		total = total + x
	}
	return total
}
