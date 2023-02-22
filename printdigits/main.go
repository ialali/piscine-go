package main

import "github.com/01-edu/z01"

func main() {
	for j := 'a'; j <= 'z'; j++ {
		z01.PrintRune(j)
	}
	z01.PrintRune('\n')
}
