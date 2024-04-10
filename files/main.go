package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("/Users/valentinemadu/go-files/files/read.txt")
	if err != nil {
		fmt.Errorf("%s", err)
	}

	if err = os.WriteFile("/Users/valentinemadu/go-files/files/write.txt", data, 0644); err != nil {
		fmt.Println(err)
	}

}
