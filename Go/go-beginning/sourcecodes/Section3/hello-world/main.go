package main

import "fmt"

func main() {
	sayHelloWorld("Hello, world again!")
}

func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}