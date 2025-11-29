package textprocessor

// handleVowels converts "a" to "an" if next word is a begins with a vowel
func HandleVowels(tokens []string) []string {
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