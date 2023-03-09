package piscine

func Abort(a, b, c, d, e int) int {
	sum := a + b + c + d + e
	max := a
	min := a

	// Find the maximum and minimum values among the arguments
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	if d > max {
		max = d
	}
	if e > max {
		max = e
	}

	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	if d < min {
		min = d
	}
	if e < min {
		min = e
	}
	return (sum - max - min) / 3
}
