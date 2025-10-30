package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func main() {
	fmt.Println("Learning CRUD...")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error while getting GET response", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error while getting GET response", res.StatusCode)
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while reading response body", err)
		return
	}
	fmt.Println("Data: ", string(data))

	var todo Todo
	err = json.Unmarshal(data, &todo)
	if err != nil {
		fmt.Println("Error while unmarshalling data", err)
		return
	}
	fmt.Println("Todo : ", todo)

	// or we can do this following decoding instead of unmarshalling     decoding = ioutil.readall + unmarshall means while decoding we dont need to readall our api response
	// err = json.NewDecoder(res.Body).Decode(&todo)
	// if err != nil {
	// 	fmt.Println("Error while decoding data", err)
	// 	return
	// }
	// fmt.Println("Todo : ", todo)

	fmt.Println("Todo Title : ", todo.Title)
	fmt.Println("Todo Completed : ", todo.Completed)
}