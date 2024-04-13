package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// jsonStream represents our raw JSON.
const jsonStream = `
	{"Name": "Ed", "Text": "Knock knock."}
	{"Name": "Sam", "Text": "Who's there?"}
`

type Message struct {
	Name string
	Text string
}

func main() {
	reader := strings.NewReader(jsonStream)

	dec := json.NewDecoder(reader)
	msgs := make(chan Message, 1)
	errs := make(chan error, 1)
	go func() {
		defer close(msgs)
		defer close(errs)
		for {
			var m Message
			if err := dec.Decode(&m); err == io.EOF {
				break
			} else if err != nil {
				errs <- err
				return
			}
			msgs <- m
		}
	}()
	for m := range msgs {
		fmt.Printf("%+v\n", m)
	}
	if err := <-errs; err != nil {
		fmt.Println("stream error: ", err)
	}
}
