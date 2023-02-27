package piscine

func IterativeFactorial(nb int) int {
	if nb >= 0 && nb <= 20 {
		f := 1
		for i := 1; i <= 20; i++ {
			f = f * i
		}
		return (f)
	} else {
		return 0
	}
}
