package main

import (
	"encoding/json"
	"fmt"
)

type Record struct {
	Name string `json:"user_name"`
	User string `json:"user"`
	ID   int    `json:"identification"`
	Age  int    `json:"-"`
}

func main() {
	rec := Record{
		Name: "John Doak",
		User: "jdoak",
		ID:   23,
	}

	b, err := json.Marshal(rec)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}
