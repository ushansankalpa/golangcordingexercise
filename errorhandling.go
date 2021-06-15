package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Failed to open file at path", pErr.Path)
			return
		}
		fmt.Println("Generic error", err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }
