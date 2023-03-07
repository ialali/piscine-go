package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename as an argument")
		return
	}

	filename := os.Args[1]
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %s", filename, err)
		return
	}

	fmt.Print(string(content))
}
