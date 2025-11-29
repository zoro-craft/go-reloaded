package textprocessor

import "strings"

// handleQuotes ensures proper spacing around single quotes
func HandleQuotes(tokens []string) []string {
	text := strings.Join(tokens, " ")
	result := ""
	quoteContent := ""
	runes := []rune(text)

	for i := 0; i < len(runes); i++ {
		if runes[i] != '\'' {
			result += string(runes[i])
			continue
		}

		if isPartOfWord(runes, i) {
			result += string(runes[i])
			continue
		}

		closingQuoteIndex := findClosingQuote(runes, i)
		if closingQuoteIndex < 0 {
			result += string(runes[i])
			continue
		}

		quoteContent = extractQuoteContent(runes, i, closingQuoteIndex)
		result += "'" + quoteContent + "'"

		i = closingQuoteIndex
	}

	return strings.Fields(result)
}