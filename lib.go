package main

import "strconv"

func handleHexAndBin(tokens []string) []string {
	// Iterate through all tokens
	for i := 0; i < len(tokens); i++ {
		currentToken := tokens[i]
		// Check if current token is a hexadecimal conversion marker
		if currentToken == "(hex)" || currentToken == "(bin)" {
			// Get the index of the previous token
			previousTokenIndex := i - 1
			// Ensure a previous token exists
			if previousTokenIndex >= 0 {
				previousToken := tokens[previousTokenIndex]
				var num int64
				var err error
				if currentToken == "(hex)" {
					// Convert hexadecimal string to decimal integer
					num, err = strconv.ParseInt(previousToken, 16, 64)
					// Skip conversion if the string is not valid hexadecimal
					if err != nil {
						continue
					}
				} else if currentToken == "(bin)" {
					// Convert binary string to decimal integer
					num, err = strconv.ParseInt(previousToken, 2, 64)
					// Skip conversion if the string is not valid binary
					if err != nil {
						continue
					}
				}
				// Convert decimal integer back to string
				newPreviousToken := strconv.FormatInt(num, 10)
				// Replace the hexadecimal or binary string with decimal string
				tokens[previousTokenIndex] = newPreviousToken
				// Remove the (hex) or (bin) markers from the slice
				tokens = append(tokens[:i], tokens[i+1:]...)
				// Decrement index since we removed an element
				i--
			}
		}
	}
	return tokens
}
