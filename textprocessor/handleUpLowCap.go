package textprocessor

// handleUpLowCap changes the case of the previous word
func HandleUpLowCap(tokens []string) []string {
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
