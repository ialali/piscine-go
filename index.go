package piscine

func Index(s string, toFind string) int {
	n := len(s)
	m := len(toFind)

	if m == 0 {
		return 0
	}

	if n < m {
		return -1
	}

	for i := 0; i <= n-m; i++ {
		if s[i:i+m] == toFind {
			return i
		}
	}

	return -1
}
