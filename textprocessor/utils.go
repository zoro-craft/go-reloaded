package textprocessor

import (
	"strconv"
	"strings"
	"unicode"
)

func capitalize(str string) string {
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
// checks if a string is a word
func isWord(str string) bool {
	for _, char := range str {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

// checks wether a rune is a letter or a digit
func isWordChar(char rune) bool {
	if unicode.IsLetter(char) || unicode.IsDigit(char) {
		return true
	}
	return false
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

// converts a string from hex or binary to decimal
func convertToDecimal(value string, markerType string) (int64, error) {
	if markerType == "(hex)" {
		return strconv.ParseInt(value, 16, 64)
	}
	return strconv.ParseInt(value, 2, 64)
}

// applies uppercase, lowercase, or capitalization
func applyCaseTransformation(word string, markerType string) string {
	switch markerType {
	case "(up)":
		return strings.ToUpper(word)
	case "(low)":
		return strings.ToLower(word)
	case "(cap)":
		return capitalize(word)
	default:
		return word
	}
}

// extracts the number from a string like "3)"
func extractCount(token string) (int, error) {
	rawNumber := strings.TrimSuffix(token, ")")
	return strconv.Atoi(rawNumber)
}

// transforms multiple previous words
func applyTransformationToMultipleWords(tokens []string, startIndex int, count int, markerType string) {
	wordsProcessed := 0

	for j := startIndex - 1; j >= 0 && wordsProcessed < count; j-- {

		if !isWord(tokens[j]) {
			continue
		}

		tokens[j] = applyCaseTransformation(tokens[j], strings.TrimSuffix(markerType, ",")+")")
		wordsProcessed++
	}
}

// checks if a quote is part of a word (like "don't")
func isPartOfWord(runes []rune, index int) bool {
	hasCharBefore := index > 0 && isWordChar(runes[index-1])
	hasCharAfter := index < len(runes)-1 && isWordChar(runes[index+1])
	return hasCharBefore && hasCharAfter
}

// finds the matching closing quote
func findClosingQuote(runes []rune, startIndex int) int {
	for j := startIndex + 1; j < len(runes); j++ {
		if runes[j] == '\'' {

			if isPartOfWord(runes, j) {
				continue
			}
			return j
		}
	}
	return -1
}

// gets the text between quotes and trims whitespace
func extractQuoteContent(runes []rune, startIndex int, endIndex int) string {
    content := string(runes[startIndex+1 : endIndex])
    return strings.TrimSpace(content)
}

// finds the index of the previous word token
func findPreviousWord(tokens []string, currentIndex int) int {
	for i := currentIndex - 1; i >= 0; i-- {
		if isWord(tokens[i]) {
			return i
		}
	}
	return -1
}

// finds the index of the next word token
func findNextWord(tokens []string, currentIndex int) int {
	for i := currentIndex + 1; i < len(tokens); i++ {
		if isWord(tokens[i]) {
			return i
		}
	}
	return -1
}

// removes a token at the specified index
func removeToken(tokens []string, index int) []string {
	return append(tokens[:index], tokens[index+1:]...)
}

// containsRune checks if a rune slice contains a specific rune
func containsRune(slice []rune, target rune) bool {
	for _, r := range slice {
		if r == target {
			return true
		}
	}
	return false
}
