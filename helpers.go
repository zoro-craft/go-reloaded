package main

import (
	"unicode"
)

func Capitalize(str string) string {
	if len(str) == 0 {
		return ""
	}
	chars := []rune(str)
	chars[0] = unicode.ToUpper(chars[0])
	for i := 1; i < len(chars); i++ {
		char := chars[i]
		chars[i] = unicode.ToLower(char)
	}
	newStr := string(chars)
	return newStr
}
