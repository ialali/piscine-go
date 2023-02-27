package main

import (
	"fmt"
	"piscine"
)

func IterativeFactorial(nb int) int {
	if nb <= 0 && nb < 20 {
		f := 1
		for int := 1; int <= 20; int++ {
			f = f * int
		}
		return (f)
	} else {
		return (0)
	}
}

func main() {
	arg := 21
	fmt.Println(piscine.IterativeFactorial(arg))
}
