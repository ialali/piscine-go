package piscine

func ToUpper(s string) string {
	result := ""
	for _, value := range s {
		if value >= 97 && value <= 122 {
			value -= 32
		}
		result += string(value)
	}
	return result
}
