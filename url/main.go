package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL")
	myUrl := "https://example.com:8080/path/to/resource?key1=value1&key2=value2"
	fmt.Printf("Type of URL: %T\n", myUrl)

	parsedURL, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println("Can't parse URL", err)
		return
	}
	fmt.Printf("Type of URL: %T\n", parsedURL)

	fmt.Println("Scheme: ", parsedURL.Scheme)
	fmt.Println("Host: ", parsedURL.Host)
	fmt.Println("Path: ", parsedURL.Path)
	fmt.Println("RawQuery", parsedURL.RawQuery)

	//Modifying URL components
	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "username=iambhavesh"

	//constructing a URL string from a URL object
	newUrl := parsedURL.String()

	fmt.Println("new URL: ", newUrl)
}

//OUTPUT:=

// Learning URL
// Type of URL: string
// Type of URL: *url.URL
// Scheme:  https
// Host:  example.com:8080
// Path:  /path/to/resource
// RawQuery key1=value1&key2=value2
// new URL:  https://example.com:8080/newPath?username=iambhavesh
