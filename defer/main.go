package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Starting")
	data := add(5, 6)
	defer fmt.Println("Data is ", data)
	defer fmt.Println("Middle")
	fmt.Println("End")

	//output = Starting
	// End
	// Middle
	// Data is  11
}
