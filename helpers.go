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

func isWord(str string) bool {
	for _, char := range str {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func isWordChar(char rune) bool {
	if unicode.IsLetter(char) || unicode.IsDigit(char) {
		return true
	}
	return false
}

func isPunctuation(char rune) bool {
	punctuations := []rune{'.', ',', '!', '?', ':', ';'}
	isPunctuation := false
	for _, punctuation := range punctuations {
		if char == punctuation {
			return true
		}
	}
	return isPunctuation
}

func doesBeginWithVowel(str string) bool {
	strRunes := []rune(str)
	switch strRunes[0] {
	case 'a', 'e', 'i', 'o', 'u', 'h':
		return true
	case 'A', 'E', 'I', 'O', 'U', 'H':
		return true
	default:
		return false
	}
}
