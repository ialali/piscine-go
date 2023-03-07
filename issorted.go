package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 1; i < len(a); i++ {
		cmp := f(a[i-1], a[i])
		if cmp > 0 {
			return false
		}
		if cmp == 0 && i > 1 && f(a[i-2], a[i-1]) != 0 {
			return false
		}
	}
	return true
}
