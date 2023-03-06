package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	slice := []int{}
	for i := min; i < max; i++ {
		slice = append(slice, i)
	}
	return slice
}
