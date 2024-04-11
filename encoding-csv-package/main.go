package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type record []string

func (r record) validate() error {
	if len(r) != 2 {
		return fmt.Errorf("data format is incorrect")
	}
	return nil
}

func (r record) first() string {
	return r[0]
}

func (r record) last() string {
	return r[1]
}

func readRecs() ([]record, error) {
	file, err := os.Open("data.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2
	reader.TrimLeadingSpace = true

	var recs []record
	for {
		data, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		rec := record(data)
		recs = append(recs, rec)
	}
	return recs, nil
}

const fakeContent = `
John,Doak
Sarah, Murphy
David, Justice
`

func main() {
	// Write our content that we will read to disk.
	file, err := os.OpenFile("data.csv", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	if _, err := file.Write([]byte(fakeContent)); err != nil {
		panic(err)
	}

	recs, err := readRecs()
	if err != nil {
		panic(err)
	}
	for _, rec := range recs {
		fmt.Println(rec.first())
	}
}
