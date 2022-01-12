package strs

import "strings"

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

func CamelToSnake(str string) string {
	var slice []string
	start := 0
	for end, current := range str {
		if end+1 < len(str) {
			next := str[end+1]
			if current < 'A' || current > 'Z' {
				if next >= 'A' && next <= 'Z' {
					// current is not upper, and next is upper
					slice = append(slice, str[start:end+1])
					start = end + 1
				}
			} else if end+2 < len(str) && (next >= 'A' && next <= 'Z') {
				if next2 := str[end+2]; next2 < 'A' || next2 > 'Z' {
					// current is upper, next is upper, and next2 is not upper
					slice = append(slice, str[start:end+1])
					start = end + 1
				}
			}
		} else {
			slice = append(slice, str[start:end+1])
		}
	}
	return strings.ToLower(strings.Join(slice, "_"))
}

func SnakeToCamel(column string) string {
	var parts []string
	for _, part := range strings.Split(column, "_") {
		parts = append(parts, FirstLetterToUpper(part))
	}
	return strings.Join(parts, "")
}
