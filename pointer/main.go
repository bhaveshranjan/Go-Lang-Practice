package main

import "fmt"

func modifyvalueByReference(num *int) {
	*num = *num + 20
}

func main() {

	var num int
	num = 2

	var ptr *int
	ptr = &num

	//fmt.Println("Num has Value: ", num)
	fmt.Println("Pointer Contains: ", ptr)
	fmt.Println("Data Contains through Pointer: ", *ptr)

	var pointer *int
	if pointer == nil {
		fmt.Println("Pointer is not assigned")
		fmt.Println(pointer)
	}

	value := 5
	modifyvalueByReference(&value)
	fmt.Println("Value Contains: ", value)

}
