package piscine

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}

	// Sort the array of integers in ascending order
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	// Calculate the median
	median := arr[2]

	return median
}
