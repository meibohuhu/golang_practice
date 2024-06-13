package staff

import (
	"fmt"
	"log"
)

// OverpaidLimit is exported, and thus visible outside of the package
var OverpaidLimit = 75000

// underPaidLimit is not exported, and thus not visible outside of the package
var underPaidLimit = 30000

// Employee holds an employee
type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}

// Office is used to initialize this package with data
type Office struct {
	AllStaff []Employee
}

// All simply returns the slice of employees used to initialize this package
func (e *Office) All() []Employee {
	return e.AllStaff
}

// Overpaid returns subset of all staff who are paid >= OverpaidLimit
func (e *Office) Overpaid() []Employee {
	var overpaid []Employee
	for _, x := range e.AllStaff {
		if x.Salary >= OverpaidLimit {
			overpaid = append(overpaid, x)
		}
	}
	return overpaid
}

// Underpaid returns subset of all staff who are paid <= underPaidLimit
func (e *Office) Underpaid() []Employee {
	var underpaid []Employee
	for _, x := range e.AllStaff {
		if x.Salary < underPaidLimit {
			underpaid = append(underpaid, x)
		}
	}
	return underpaid
}

// notVisible is not exported, thus is only available in this package
func (e *Office) notVisible() {
	fmt.Println("Hello, world")
}

// myFunction is not exported, and does not have a receiver
// It can only be used in this package, and does not have access to the values
// stored in the receiver variable (of type *Office)
func myFunction() {
	log.Println("I am a function")
}
