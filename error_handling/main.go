package main

import "fmt"

// func divide(a, b float64) float64 {
// 	return a / b
// }

// func divide(a, b float64) float64 {
// 	if b == 0 {
// 		return 0
// 	}
// 	return a / b
// }

// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("denominator must not be Zero")
// 	}
// 	return a / b, nil
// }

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator must not be Zero")
	}
	return a / b, nil
}

func main() {
	// fmt.Println("Started Error Handling in the Functions")

	// ans, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println("Error Handling")
	// }
	// //ans := divide(10, 0)
	// fmt.Println("Division of two numbers is", ans)

	data, _ := divide(10, 0)
	fmt.Println(data)

	// fmt.Println("Started Error Handling in the Functions")
	// ans, _ := divide(10, 0)
	// fmt.Println("Division of two numbers is", ans)
}
