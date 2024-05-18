package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	var bhavesh Person
	fmt.Println("Bhavesh Person :", bhavesh)
	bhavesh.FirstName = "Bhavesh"
	bhavesh.LastName = "Ranjan"
	bhavesh.Age = 24

	fmt.Println("Bhavesh Person :", bhavesh)

	person1 := Person{
		FirstName: "Devesh",
		LastName:  "Ranjan",
		Age:       22,
	}
	fmt.Println("Person 1", person1)

	var person2 = new(Person) //new keyword person2 is a kind of pointer
	person2.FirstName = "Rajesh"
	person2.LastName = "Kumar"
	person2.Age = 26

	fmt.Println("Person2 :", person2)
}
