package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Learning CRUD...")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting : ", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting Response :", res.Status)
		return
	}
	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading: ", err)
	// 	return
	// }
	// fmt.Println("Data: ", string(data))

	// Output :=
	// {
	// 	"userId": 10,
	// 	"id": 199,
	// 	"title": "numquam repellendus a magnam",
	// 	"completed": true
	// }

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding : ", err)
		return
	}
	fmt.Println("Todo: ", todo)
}

//Output:=

// Learning CRUD...
// Todo:  {1 1 delectus aut autem false}