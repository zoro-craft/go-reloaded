package textprocessor

// handlePunctuation ensures proper spacing around punctuation marks
func HandlePunctuation(tokens []string) []string {
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