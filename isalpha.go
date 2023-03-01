package piscine

func IsAlpha(s string) bool {
	bool := true

	for _, value := range s {
		if value >= 'a' && value <= 'z' || value >= 'A' && value <= 'Z' || value >= '0' && value <= '9' {
			bool = bool
		} else {
			bool = false
		}
	}
	return bool
}
