package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

// filePath is the path to file we are going to read and write to.
const filePath = "users.txt"

// init is a neat feature in Go. Every package can have init() functions and it is
// the only type of function whose name can be duplicated. init() is executed when
// the package loads, usually to initialize some variables or do other setup. In this
// case we use it to write some content to a file representing our users.
func init() {
	content := []byte("jdoak:0\nsmurphy:1\ndjustice:2")

	if err := os.WriteFile(filePath, content, 0644); err != nil {
		panic(err)
	}
}

// User represents our user data.
type User struct {
	Name string
	ID   int
	Err  error
}

// String implememnts fmt.Stringer. It will output the data as "user:id", such as "jdoak:0".
func (u User) String() string {
	return fmt.Sprintf("%s:%d", u.Name, u.ID)
}

// getUser takes a string that should be formatted as [user]:[id], such as "jdoak:0" and returns
// a User object.
func getUser(s string) (User, error) {
	sp := strings.Split(s, ":")
	if len(sp) != 2 {
		return User{}, fmt.Errorf("record(%s) was not in the correct format", s)
	}
	id, err := strconv.Atoi(sp[1])
	if err != nil {
		return User{}, fmt.Errorf("record(%s) had non-numeric ID", s)
	}
	return User{Name: strings.TrimSpace(sp[0]), ID: id}, nil
}

func decodeUsers(ctx context.Context, r io.Reader) chan User {
	ch := make(chan User, 1)
	// Spin of goroutine off to feed the channel we will return.
	go func() {
		// Close our channel on exit, signaling we are done.
		defer close(ch)

		// Wrap a Scanner around our reader so we can read each line of content.
		scanner := bufio.NewScanner(r)
		for scanner.Scan() { // Scan until nothing to scan.
			if ctx.Err() != nil { // Context was cancelled, return error.
				ch <- User{Err: ctx.Err()}
				return
			}
			// Turn the line of text into a User object.
			u, err := getUser(scanner.Text())
			if err != nil { // line was in incorrect format, return an error.
				u.Err = err
				ch <- u
				return
			}
			// Everything was fine, return a user record.
			ch <- u
		}
	}()

	// Returns the channel we will read off of.
	return ch
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Open the file we created with init().
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Start decoding the file one line at a time.
	ch := decodeUsers(ctx, f)

	// Read each line of output and write the record to the screen.
	for u := range ch {
		if u.Err != nil {
			panic(u.Err)
		}
		fmt.Println(u)
	}
}
