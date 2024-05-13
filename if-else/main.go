package main

import "fmt"

func main() {
	z := 5
	if z > 5 {
		fmt.Println("10 is greater than 5")
	} else if z < 5 {
		fmt.Println("5 is greater than 10")
	} else {
		fmt.Println("Both the variables are equal")
	}
	x := 7
	y := 10
	if x > 5 && y > 5 {
		fmt.Println("hey how are you ?")
	} else {
		fmt.Println("Learn programming with me")
	}
}
