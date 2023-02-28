package piscine

func LastRune(s string) rune {
	arrayStr := []rune(s)
	return []rune(arrayStr)[len(arrayStr)-1]

}
