package main

import (
	"encoding/json"
	"fmt"
)

//Person a
type Person struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func main() {
	in := `{"firstName":"Jhon","lastName":"Dow"}`
	bytes := []byte(in)

	var p Person
	err := json.Unmarshal(bytes, &p)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", p)
}
