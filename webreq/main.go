package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning web service")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting GET response ", err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("Type of response: %T\n", res)

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response", err)
		return
	}
	fmt.Println("response:", string(data))
}

// Output:-
// Learning web service
// Type of response: *http.Response
// response: {
//   "userId": 1,
//   "id": 1,
//   "title": "delectus aut autem",
//   "completed": false
// }
