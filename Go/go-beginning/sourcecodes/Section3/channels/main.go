package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

var keyPressChan chan rune

func main() {
	keyPressChan = make(chan rune)

	go listenForKeyPress()

	fmt.Println("Press any key, or q to quit")
	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}

		keyPressChan <- char
	}
}

// listenForKeyPress is called using the "go" keyword, so it runs as a goroutine while the
// calling function (main) continues to execute
func listenForKeyPress() {
	for {
		key := <-keyPressChan    // listen the value of whatever is passed in
		fmt.Println("You pressed", string(key))
	}
}
