package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello World !")
	time.Sleep(2000 * time.Millisecond) // Simulating some work
	fmt.Println("Say hello function ended")
}

func sayHi() {
	fmt.Println("Hi Bhavesh :)")
}

func main() {
	fmt.Println("learning goroutines")
	go sayHello()
	go sayHi()

	//wait for a moment to allow the go routine to finish

	time.Sleep(3000 * time.Millisecond)
}

// Output :-

// learning goroutines
// Hi Bhavesh :)
// Hello World !
// Say hello function ended
