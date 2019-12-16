package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func main() {
	// fmt.Println("Test")
	// fmt.Println("Hello World")
	var JsonString = `{"name": "Dzakiy", "age": "26"}`
	var JsonData = []byte(JsonString)

	var data User
	var err = json.Unmarshal(JsonData, &data)

	if err != nil {
		fmt.Println("Error Unmarhasilling Json " + err.Error())
		return
	}

	fmt.Println(data)
	fmt.Println("My name is " + data.Name)
	fmt.Println("My age is " + data.Age)

	// json.Unmarshal(JsonData)
	// fmt.Println(JsonData)
}
