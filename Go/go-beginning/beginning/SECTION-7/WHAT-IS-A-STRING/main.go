package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println()
	name := "Hello world"
	fmt.Println("String:", name)
	fmt.Println()

	fmt.Println("Bytes")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%x ", string(name[i])) // still bytes
		fmt.Println(string(name[i]))       // string
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("Index\tRune\tString")
	for x, y := range name {
		fmt.Println(x, "\t", y, "\t", string(y))
	}
	fmt.Println()

	fmt.Println()
	fmt.Println("Three ways to concatenate strings")

	h := "Hello, "
	w := "world."

	// using + not very efficient
	myString := h + w
	fmt.Println(myString)
	fmt.Println()

	// using fmt - more efficient
	myString = fmt.Sprintf("%s%s", h, w)
	fmt.Println(myString)
	fmt.Println()

	// // using stringbuilder - very efficient

	var sb strings.Builder
	sb.WriteString(h)
	sb.WriteString(w)
	fmt.Println(sb.String())
	fmt.Println(myString[10:13])

	fmt.Println()
	// indexes start at 0
	//                       1         2         3
	//             01234567890123456789012345678901234
	courseName := "Learn Go for Beginners Crash Course"
	fmt.Println(courseName[0])
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
