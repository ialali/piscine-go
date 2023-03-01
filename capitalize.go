package piscine

func Capitalize(s string) string {
	output := ""
	capitalizeNext := true
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			if capitalizeNext {
				output += ToUpper(string(char))
			} else {
				output += ToLower(string(char))
			}
			capitalizeNext = false
		} else {
			capitalizeNext = true
			output += string(char)
		}
	}
	return output
}
