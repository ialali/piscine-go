package piscine

func ToLower(s string) string {
	result := ""
	for _, value := range s {
		if value >= 65 && value <= 90 {
			value += 32
		}
		result += string(value)
	}
	return result
}
