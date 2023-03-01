package piscine

func IsNumeric(s string) bool {
	isValid := true

	for _, value := range s {
		if value >= '0' && value <= '9' {
			// the character is a digit, so isValid remains true
		} else {
			// the character is not a digit, so the string is not numeric
			isValid = false
			break
		}
	}

	return isValid
}
