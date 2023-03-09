package piscine

func Abort(a, b, c, d, e int) int {
	min := a
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

	max := a
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

	sum := a + b + c + d + e

	middleSum := sum - min - max

	return middleSum / 3
}
