package main

import "fmt"

func main() {
	fmt.Println("we are learning Array in Golang")

	var name [5]string

	// 0 1 2 3 4
	name[0] = "prince"
	name[2] = "Akash"

	fmt.Println("Names of Person is :", name)

	var numbers = [8]int{1, 2, 3, 4, 5}
	fmt.Println("Number is :", numbers)

	fmt.Println("Length of Numbers array is:", len(numbers))

	fmt.Println("value of name at 2 index is :", name[2])
}
