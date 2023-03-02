package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args
	for _, b := range a[1:] {
		for _, r := range b {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
