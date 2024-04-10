package main

import (
	"context"
	"errors"
	"io"
	"net/http"
	"os"
)

func readRemote(ctx context.Context) error {
	client := http.Client{}
	req, err := http.NewRequest("GET", "http://myserver.mydomain/myfile", nil)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)
	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	// resp contains an io.ReadCloser that we can read as a file.
	// Let's use io.ReadAll() to read the entire content to data.
	// data, err := io.ReadAll(resp.Body)

	flags := os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	f, err := os.OpenFile("path/to/file", flags, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := io.Copy(f, resp.Body); err != nil {
		return err
	}
	return errors.New("did not complete succesfully")
}

func main() {
	readRemote(context.Background())
}
