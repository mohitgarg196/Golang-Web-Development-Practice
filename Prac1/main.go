package main

import (
	"fmt"
	"net/http"
)

func printTasks(){
	fmt.Println("### My Tasks ###")
	for index, item:= range allItems{
		fmt.Println(index+1, ".", item)
	}
}

func helloHandler(writer http.ResponseWriter, request *http.Request){
	var greeting = "Hello, Golang World!"
	fmt.Fprintln(writer, greeting);
}

var allItems = []string {"Golang", "Python", "JavaScript", "C++", "Java"}
func main(){
	var tech = "Golang"
	printTasks()
	fmt.Println(tech)
	fmt.Println(allItems)
	
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":3080", nil)
}