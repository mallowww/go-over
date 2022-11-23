package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {
	data, _ := json.Marshal(&employee{101, "Ki To", "0800000", "fakemail@gmail.com"})
	fmt.Println(string(data))
}
