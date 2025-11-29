package textprocessor

import "strings"

// handleNumberedUpLowCap handles case changes for multiple previous words
func HandleNumberedUpLowCap(tokens []string) []string {
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