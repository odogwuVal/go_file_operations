package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"
)

const flags = os.O_CREATE | os.O_WRONLY | os.O_TRUNC

// User represents our user data.
type User struct {
	// Name is our user's username.
	Name string
	// ID is their unique numeric ID in the system.
	ID int
}

// String implememnts fmt.Stringer. It will output the data as "user:id", such as "jdoak:0".
func (u User) String() string {
	return fmt.Sprintf("%s:%d", u.Name, u.ID)
}

func writeUser(ctx context.Context, w io.Writer, u User) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	if _, err := w.Write([]byte(u.String())); err != nil {
		return err
	}

	return nil
}

// WriteUsers writes a list of users to 'w' with each entry separated by '\n'.
func WriteUsers(ctx context.Context, w io.Writer, users []User) error {
	for i, u := range users {
		if i != 0 {
			if _, err := w.Write([]byte("\n")); err != nil {
				return err
			}
		}
		if err := writeUser(ctx, w, u); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// filePath represents the file name we want to write to.
	const filePath = "file.txt"

	// users represents a list of user data we want to write. You will notice we don't
	// put the type User{} around every entry in []User{}. Go can infer the type from the
	// list defintion.
	var users = []User{
		{Name: "jdoak", ID: 0}, // Shorthand syntax
		{Name: "smurphy", ID: 1},
		User{Name: "djustice", ID: 2}, // Long syntax
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Open the file and get the *os.File object. *os.File implements io.Writer among
	// other interfaces.
	f, err := os.OpenFile(filePath, flags, 0644)
	if err != nil {
		panic(err)
	}

	// Write our users to the file.
	if err := WriteUsers(ctx, f, users); err != nil {
		panic(err)
	}
	f.Close() // Close the file to writing

	// Open the file and write its content to stdout to show it worked.
	rf, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer rf.Close()

	io.Copy(os.Stdout, rf)
}
