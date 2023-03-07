package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
        fmt.Println("File name missing")
    } else if len(arguments) > 2 {
		fmt.Println("Too many arguments")
	} else {
		fileName := arguments[1]
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("The mistake is: %v\n", err.Error())
		} else {
			arr := make([]byte, 14)
			file.Read(arr)
			fmt.Println(string(arr))
			file.Close()
		}
	}
}
