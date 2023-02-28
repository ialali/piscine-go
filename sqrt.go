package piscine

func Sqrt(nb int) int {
	for int := 0; int <= nb; int++ {
		if int*int == nb {
			return int
		}
	}
	return 0
}
