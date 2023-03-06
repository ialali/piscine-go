package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	value := make([]int, max-min)
	for a := range value {
		value[a] = min + a
	}
	return value
}
