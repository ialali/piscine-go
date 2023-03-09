package piscine

func Rot14(s string) string {
	rot := make([]byte, len(s))
	for i, c := range []byte(s) {
		if c >= 'a' && c <= 'z' {
			rot[i] = (c-'a'+14)%26 + 'a'
		} else if c >= 'A' && c <= 'Z' {
			rot[i] = (c-'A'+14)%26 + 'A'
		} else {
			rot[i] = c
		}
	}
	return string(rot)
}
