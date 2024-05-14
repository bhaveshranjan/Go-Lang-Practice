package main

import "fmt"

func main() {
	studentGrades := make(map[string]int)

	studentGrades["Bhavesh"] = 97
	studentGrades["Prince"] = 94
	studentGrades["Akash"] = 99
	fmt.Println("Marks of Prince :", studentGrades["Prince"])

	studentGrades["Prince"] = 97
	fmt.Println("New Marks of Prince :", studentGrades["Prince"])

	delete(studentGrades, "Prince")
	fmt.Println("Marks of Prince :", studentGrades["Prince"])

	//checking if key exists or not

	grades, exists := studentGrades["Bhavesh"]
	fmt.Println("Grades of Bhavesh: ", grades)
	fmt.Println("Bhavesh exists: ", exists)

	for index, value := range studentGrades {
		fmt.Printf("Key is %s and marks is %d\n", index, value)
	}

}
