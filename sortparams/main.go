package main

import (
	"os"
	"sort"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	sort.Strings(args)
	for _, arg := range args {
		for _, r := range arg {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
