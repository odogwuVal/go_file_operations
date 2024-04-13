package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// jsonStream represents our raw JSON.
const jsonStream = `[
	{"Name": "Ed", "Text": "Knock knock."},
	{"Name": "Sam", "Text": "Who's there?"}
]
`

type Message struct {
	Name string
	Text string
}

func main() {
	reader := strings.NewReader(jsonStream)

	dec := json.NewDecoder(reader)

	_, err := dec.Token() // Reads [
	if err != nil {
		panic(fmt.Errorf(`outter [ is missing`))
	}

	for dec.More() {
		var m Message
		// decode an array value (Message)
		err := dec.Decode(&m)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%+v\n", m)
	}

	_, err = dec.Token() // Reads ]
	if err != nil {
		panic(fmt.Errorf(`final ] is missing`))
	}
}
