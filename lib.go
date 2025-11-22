package main

import (
	"strconv"
	"strings"
)

func handleHexAndBin(tokens []string) []string {
	// Iterate through all tokens
	for i := 0; i < len(tokens); i++ {
		currentToken := tokens[i]
		// Check if current token is a hexadecimal or binary conversion marker
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
						// Remove the (hex) or (bin) markers from the slice
						tokens = append(tokens[:i], tokens[i+1:]...)
						// Decrement index since we removed an element
						i--
						continue
					}
				} else if currentToken == "(bin)" {
					// Convert binary string to decimal integer
					num, err = strconv.ParseInt(previousToken, 2, 64)
					// Skip conversion if the string is not valid binary
					if err != nil {
						// Remove the (hex) or (bin) markers from the slice
						tokens = append(tokens[:i], tokens[i+1:]...)
						// Decrement index since we removed an element
						i--
						continue
					}
				}
				// Convert decimal integer back to string
				newPreviousToken := strconv.FormatInt(num, 10)
				// Replace the hexadecimal or binary string with decimal string
				tokens[previousTokenIndex] = newPreviousToken
			}
			// Remove the (hex) or (bin) markers from the slice
			tokens = append(tokens[:i], tokens[i+1:]...)
			// Decrement index since we removed an element
			i--
		}
	}
	return tokens
}

func handleUpLowCap(tokens []string) []string {
	//iterate through all token
	for i := 0; i < len(tokens); i++ {
		currentToken := tokens[i]
		//check if current token is lowercase or uppercase or capitalize conversion marker
		if currentToken == "(up)" || currentToken == "(low)" || currentToken == "(cap)" {
			//get the index of the previous token
			previousTokenIndex := i - 1
			//ensure a previous token exists
			if previousTokenIndex >= 0 {
				previousToken := tokens[previousTokenIndex]
				var newPreviousToken string
				if currentToken == "(up)" {
					//convert to uppercase
					newPreviousToken = strings.ToUpper(previousToken)
				} else if currentToken == "(low)" {
					//convert to lower case
					newPreviousToken = strings.ToLower(previousToken)
				} else if currentToken == "(cap)" {
					//capitalize
					newPreviousToken = Capitalize(previousToken)
				}
				//replace the previous token with the new token
				tokens[previousTokenIndex] = newPreviousToken
			}
			//remove the (up) or (low) or (cap) markers from the slice
			tokens = append(tokens[:i], tokens[i+1:]...)
			//Decrement the index since we removed an element
			i--

		}
	}
	return tokens
}
