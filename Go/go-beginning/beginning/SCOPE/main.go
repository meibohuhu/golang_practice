package main

import (
	"fmt"
	"myapp/packageone"
)

var myVar = "This is a package level variable"

func main() {

	newString := packageone.PackageVar
	fmt.Println(newString)
	// variables
	var blockVar = "This is the block level variable"
	packageone.PrintMe(myVar, blockVar)
}
