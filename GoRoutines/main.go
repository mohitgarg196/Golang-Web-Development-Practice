package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, World!")
	time.Sleep(1000*time.Millisecond)
	fmt.Println("Hello, World! after 1000 milliseconds")
}

func sayHi() {
	time.Sleep(1*time.Millisecond)
	fmt.Println("Hi, Mohit Garg :-)")
}

func main() {
	fmt.Println("Learning GoRoutines")
	
	// Sequential execution
	sayHello()
	sayHi()

	// Concurrent execution
	go sayHello()
	sayHi()
	time.Sleep(1000*time.Millisecond)
}
