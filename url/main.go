package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning about URLs")

	myUrl := "http://localhost:8080/api/resouce?key1=9&key2=10"
	fmt.Printf("Type of URL: %T\n", myUrl)
	fmt.Println("URL: ", myUrl)

	parsedUrl, err := url.Parse(myUrl)
	
	if err != nil {
		fmt.Println("Error while parsing URL", err)
		return
	}
	fmt.Printf("Type of URL: %T\n", parsedUrl)
	fmt.Println("URL: ", parsedUrl)

	fmt.Println("Scheme: ", parsedUrl.Scheme)
	fmt.Println("Host: ", parsedUrl.Host)
	fmt.Println("Path: ", parsedUrl.Path)
	fmt.Println("Raw Query: ", parsedUrl.RawQuery)
}