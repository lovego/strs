package strs

func FirstLetterToUpper(s string) string {
	if len(s) > 0 && s[0] >= 'a' && s[0] <= 'z' {
		b := []byte(s)
		b[0] -= ('a' - 'A')
		return string(b)
	}
	return s
}

func FirstLetterToLower(s string) string {
	if len(s) > 0 && s[0] >= 'A' && s[0] <= 'Z' {
		b := []byte(s)
		b[0] += ('a' - 'A')
		return string(b)
	}
	return s
}
