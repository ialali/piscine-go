package piscine

import "sort"

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}
	sort.Ints(arr)
	return arr[2]
}
