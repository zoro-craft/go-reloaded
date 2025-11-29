package textprocessor

import "strconv"

// handleHexAndBin converts hexadecimal and binary numbers to decimal format
func HandleHexAndBin(tokens []string) []string {

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