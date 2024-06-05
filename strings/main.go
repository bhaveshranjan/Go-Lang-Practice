package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,orange,banana,mango"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	//output = [apple orange banana mango]

	str := "one two three four two two five"
	count := strings.Count(str, "t")
	fmt.Println("Count: ", count)

	//output = Count:  4

	str = "                 Hello,              Bhavesh         "
	trimmed := strings.TrimSpace(str)
	fmt.Println("trimmed: ", trimmed)

	// it trims or remove all the space present before sentence of end of sentence
	// but it will not remove any kind of spaces present in between of spaces

	//output = trimmed:  Hello,              Bhavesh

	str1 := "Bhavesh"
	str2 := "Ranjan"
	result := strings.Join([]string{str1, "XYZ", str2}, ", ,")
	fmt.Println("result: ", result)

	//output = result:  Bhavesh, ,Ranjan 
	// result:  Bhavesh, ,XYZ, ,Ranjan
}
