package main

import (
	"strings"
	"unicode"
)

func capitalize(str string) string {
	words := strings.Split(str, " ")
	newWords := make([]string, len(words))
	for i, word := range words {
		if word == "" {
			newWords[i] = ""
			continue
		}
		newWord := ""
		for j, char := range word {
			if j == 0 {
				newWord += string(unicode.ToUpper(char))
			} else if j > 0 {
				newWord += string(unicode.ToLower(char))
			}
		}
		newWords[i] = newWord
	}
	newStr := strings.Join(newWords, " ")
	return newStr
}
