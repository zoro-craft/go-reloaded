package main

import (
	"strconv"
	"strings"
)

// handleHexAndBin converts hexadecimal and binary numbers to decimal format
func handleHexAndBin(tokens []string) []string {

	for i := 0; i < len(tokens); i++ {
		currentToken := tokens[i]

		if currentToken != "(hex)" && currentToken != "(bin)" {
			continue
		}

		// find the word to convert (skip non word tokens)
		targetIndex := findPreviousWord(tokens, i)
		if targetIndex < 0 {
			tokens = removeToken(tokens, i)
			i--
			continue
		}

		wordToConvert := tokens[targetIndex]
		convertedNumber, err := convertToDecimal(wordToConvert, currentToken)

		if err != nil {
			tokens = removeToken(tokens, i)
			i--
			continue
		}

		tokens[targetIndex] = strconv.FormatInt(convertedNumber, 10)

		tokens = removeToken(tokens, i)
		i--
	}

	return tokens
}

// handleUpLowCap changes the case of the previous word
func handleUpLowCap(tokens []string) []string {
	for i := 0; i < len(tokens); i++ {
		currentToken := tokens[i]

		if currentToken != "(up)" && currentToken != "(low)" && currentToken != "(cap)" {
			continue
		}

		// find the word to modify
		targetIndex := findPreviousWord(tokens, i)
		if targetIndex < 0 {
			tokens = removeToken(tokens, i)
			i--
			continue
		}

		tokens[targetIndex] = applyCaseTransformation(tokens[targetIndex], currentToken)

		tokens = removeToken(tokens, i)
		i--
	}

	return tokens
}



// handleNumberedUpLowCap handles case changes for multiple previous words
func handleNumberedUpLowCap(tokens []string) []string {
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		if token != "(up," && token != "(low," && token != "(cap," {
			continue
		}

		if i+1 >= len(tokens) || !strings.HasSuffix(tokens[i+1], ")") {
			continue
		}

		count, err := extractCount(tokens[i+1])
		if err != nil {
			continue
		}

		applyTransformationToMultipleWords(tokens, i, count, token)

		tokens = append(tokens[:i], tokens[i+2:]...)
		i--
	}

	return tokens
}

// handlePunctuation ensures proper spacing around punctuation marks
func handlePunctuation(tokens []string) []string {
	punctuations := []rune{'.', ',', '!', '?', ':', ';'}

	// first process: move leading punctuation to previous token
	for i := 0; i < len(tokens); i++ {
		currentToken := tokens[i]
		currentRunes := []rune(currentToken)

		for j := 0; j < len(currentRunes); j++ {
			char := currentRunes[j]

			if !containsRune(punctuations, char) {
				break
			}

			if i > 0 {
				tokens[i-1] += string(char)
			} else {
				// no previous token, create one at the beginning
				tokens = append([]string{string(char)}, tokens...)
				i++
			}

			currentRunes = append(currentRunes[:j], currentRunes[j+1:]...)
			tokens[i] = string(currentRunes)
			j--
		}

		if tokens[i] == "" {
			tokens = removeToken(tokens, i)
			i--
		}
	}

	// second process: handle punctuation in the middle of tokens
	for i := 0; i < len(tokens); i++ {
		currentRunes := []rune(tokens[i])

		for j := 0; j < len(currentRunes); j++ {
			char := currentRunes[j]

			if !containsRune(punctuations, char) && j > 0 && containsRune(punctuations, currentRunes[j-1]) {
				// split token at this position
				firstPart := string(currentRunes[:j])
				secondPart := string(currentRunes[j:])

				tokens = append(tokens[:i], append([]string{firstPart, secondPart}, tokens[i+1:]...)...)
				break
			}
		}
	}

	return tokens
}


// handleQuotes ensures proper spacing around single quotes
func handleQuotes(tokens []string) []string {
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


// handleVowels converts "a" to "an" if next word is a begins with a vowel
func handleVowels(tokens []string) []string {
	for i := 0; i < len(tokens); i++ {
		if tokens[i] != "a" && tokens[i] != "A" {
			continue
		}

		// find the next actual word (skip non word tokens)
		nextWordIndex := findNextWord(tokens, i)
		if nextWordIndex < 0 {
			continue
		}

		if doesBeginWithVowel(tokens[nextWordIndex]) {
			if tokens[i] == "a" {
				tokens[i] = "an"
			} else {
				tokens[i] = "An"
			}
		}
	}

	return tokens
}

