package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"sort"
)

// record represents a record containing first and last names that are
// stored in a csv.
type record []string

// validate validates if the csv line was had the correct number of entries.
func (r record) validate() error {
	if len(r) != 2 {
		return fmt.Errorf("data format is incorrect")
	}
	return nil
}

// first returns the record's first name.
func (r record) first() string {
	return r[0]
}

// last returns the record's last name.
func (r record) last() string {
	return r[1]
}

func (r record) csv() []byte {
	b := bytes.Buffer{}
	for _, field := range r {
		b.WriteString(field + ",")
	}
	b.WriteString("\n")
	return b.Bytes()
}

func writeRecs(recs []record) error {
	file, err := os.OpenFile("data-sorted.csv", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Sort by last name
	sort.Slice(recs, func(i, j int) bool { return recs[i].last() < recs[j].last() })

	w := csv.NewWriter(file)
	defer w.Flush()

	for _, rec := range recs {
		if err := w.Write(rec); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	recs := []record{
		{"John", "Doak"},
		{"Sarah", "Murphy"},
		{"David", "Justice"},
	}

	if err := writeRecs(recs); err != nil {
		panic(err)
	}

	data, err := os.ReadFile("data-sorted.csv")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", data)
}
