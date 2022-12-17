package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID   int
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	u := User{
		ID: 1, Name: "okiebro", Age: 15,
	}

	b, err := json.Marshal(u)
	fmt.Printf("byte: %T \n", b)
	fmt.Printf("byte: %s \n", b)
	fmt.Println(err)
}
