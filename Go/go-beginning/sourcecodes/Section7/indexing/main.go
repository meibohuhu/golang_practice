package main

import "fmt"

func main() {
	// indexes start at 0
	//                       1         2         3
	//             01234567890123456789012345678901234
	courseName := "Learn Go for Beginners Crash Course"

	fmt.Println(string(courseName[0]))
	fmt.Println(string(courseName[6]))


	for i := 0; i <= 21; i++ {
		fmt.Print(string(courseName[i]))
	}
	fmt.Println()

	for i := 13; i <= 21; i++ {
		fmt.Print(string(courseName[i]))
	}
	fmt.Println()

}