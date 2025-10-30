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
		fmt.Println("Error while getting GET response", err)
		return
	}

	defer res.Body.Close()
	fmt.Printf("Type of Response: %T\n", res)
	fmt.Println("Response: ", res)

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while reading response body", err)
		return
	}
	fmt.Println("Response Body: ", string(data))
}