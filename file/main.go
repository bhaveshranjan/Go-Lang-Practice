package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	/*
		file, err := os.Create("example.txt")
		if err != nil {
			fmt.Println("Error while creating file: ", err)
			return
		}
		defer file.Close()

		content := "hello Bhavesh"
		byte, errors := io.WriteString(file, content+"\n")
		fmt.Println(byte)
		if errors != nil {
			fmt.Println("Error while writing file: ", errors)
			return
		}

		fmt.Println("successfully created file")
	*/
	/*
		file, err := os.Open("example.txt")
		if err != nil {
			fmt.Println("Error while opening file ", err)
			return
		}
		defer file.Close()

		//Create a buffer to read the file content
		buffer := make([]byte, 1024)

		// Read the file content into the buffer
		for {
			n, err := file.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Error while reading file")

			}
			// process to read content
			fmt.Println(string(buffer[:n]))
		}
	*/

	// READ the entire file into a byte slice
	content, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error while reading file ", err)
		return
	}
	fmt.Println(string(content))
}
