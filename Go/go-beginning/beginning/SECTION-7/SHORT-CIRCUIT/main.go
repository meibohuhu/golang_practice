package main

import (
	"errors"
	"fmt"
)

func main() {
	x := 12
	y := 6
	c, err := divideTwoNumbers(x, y)
	if err != nil {
		fmt.Print(err)
	} else {
		if c == 2 {
			fmt.Println("We found 2")
		}
	}

}

func divideTwoNumbers(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Invalid y")
	} else {
		return x / y, nil
	}
}
