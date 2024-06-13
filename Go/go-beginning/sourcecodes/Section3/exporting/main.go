package main

import (
	"log"
	"myapp/staff"
)

// employees is the data we'll use. Note that it's type is taken from the (exported) type
// in the staff package
var employees = []staff.Employee {
	{FirstName: "John", LastName: "Smith", Salary: 30000, FullTime: true},
	{FirstName: "Sally", LastName: "Jones", Salary: 50000, FullTime: true},
	{FirstName: "Mark", LastName: "Smithers", Salary: 60000, FullTime: true},
	{FirstName: "Allan", LastName: "Anderson", Salary: 15000, FullTime: false},
	{FirstName: "Margaret", LastName: "Carter", Salary: 100000, FullTime: true},
}

func main() {
	myStaff := staff.Office {
		AllStaff: employees,
	}

	// staff.OverpaidLimit = 20000
	
	overPaid := myStaff.Overpaid()

	log.Println(overPaid)

	log.Println("Overpaid limit is", staff.OverpaidLimit)
}