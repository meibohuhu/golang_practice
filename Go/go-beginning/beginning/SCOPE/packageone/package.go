package packageone

import "fmt"

var privateVar = "This is private variable only access in package"
var PackageVar = "This is a package level variable in packageone"

func PrintMe(s1, s2 string) { // export must be uppercase
	fmt.Print(s1, s2, PackageVar)
}
