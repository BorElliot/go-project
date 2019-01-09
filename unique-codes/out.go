package main

import (
	"fmt"
)

func myFunc() {
	// Doing something concurrently
	for i := 0; i < 10; i++ {
		// time.Sleep(time.Millisecond * 100)
		fmt.Println(i)
	}
}

func main() {
	go myFunc() // myFunc will run concurrently to the main implicit goroutine

	// Cheat and wait for user input to wait for the completion of the concurrent goroutine
	var input string
	fmt.Scanln(&input)
}
