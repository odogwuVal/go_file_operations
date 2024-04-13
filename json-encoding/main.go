package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type Message struct {
	Name string
	Text string
}

// encodeMsgs takes an input channel of messages and writes the JSON encoding
// of those message to output. The returned error channel is closed when
// processing has been finished.
func encodeMsgs(in chan Message, output io.Writer) chan error {
	errs := make(chan error, 1)
	go func() {
		defer close(errs)
		enc := json.NewEncoder(output)
		for msg := range in {
			if err := enc.Encode(msg); err != nil {
				errs <- err
				return
			}
		}
	}()
	return errs
}

func main() {
	msgs := []Message{
		{Name: "Ed", Text: "Knock knock."},
		{Name: "Sam", Text: "Who's there?"},
	}

	in := make(chan Message, 1)
	out := bytes.NewBuffer([]byte{}) // Could also use a file or HTML request body

	// Start our encoder.
	errs := encodeMsgs(in, out)

	// Send messages to be encoded onto to the "out" buffer above. If the encoder has
	// an error, stop.
	for _, msg := range msgs {
		select {
		case in <- msg:
		case err := <-errs:
			panic(err)
		}
	}
	close(in) // Signal we are done with the encoder

	// Wait until we are done processing and check for errors.
	if err := <-errs; err != nil {
		panic(err)
	}

	// Read the JSON output that was written.
	b, err := io.ReadAll(out)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}
