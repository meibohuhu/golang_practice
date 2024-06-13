package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	// whatToSay := "Hello World"
	// sayHelloWorld(whatToSay)

	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)

	reader := bufio.NewReader(os.Stdin)
	for { // no other parameter
		fmt.Print("-> ")

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1) // for mac users, remove last remove

		if userInput == "quit" {
			break
		} else {
			fmt.Println(doctor.Response(userInput))
		}
	}

}

// func sayHelloWorld(whatToSay string) {
// 	fmt.Println(whatToSay)
// }
