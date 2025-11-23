package main

import (
	"strconv"
	"strings"
)

func handleHexAndBin(tokens []string) []string {
	// iterate through all tokens
	for i := 0; i < len(tokens); i++ {
		currentToken := tokens[i]
		// check if current token is a hexadecimal or binary conversion marker
		if currentToken == "(hex)" || currentToken == "(bin)" {
			// get the index of the previous token
			indexToProcess := i - 1
			// go back the previous token if its not a word
			for indexToProcess >= 0 && !isWord(tokens[indexToProcess]) {
				indexToProcess--
			}
			if indexToProcess >= 0 {
				previousToken := tokens[indexToProcess]
				var num int64
				var err error
				if currentToken == "(hex)" {
					// convert hexadecimal string to decimal integer
					num, err = strconv.ParseInt(previousToken, 16, 64)
					// skip conversion if the string is not valid hexadecimal
					if err != nil {
						// remove the (hex) or (bin) markers from the slice
						tokens = append(tokens[:i], tokens[i+1:]...)
						// decrement index since we removed an element
						i--
						continue
					}
				} else if currentToken == "(bin)" {
					// convert binary string to decimal integer
					num, err = strconv.ParseInt(previousToken, 2, 64)
					// skip conversion if the string is not valid binary
					if err != nil {
						// remove the (hex) or (bin) markers from the slice
						tokens = append(tokens[:i], tokens[i+1:]...)
						// decrement index since we removed an element
						i--
						continue
					}
				}
				// convert decimal integer back to string
				newPreviousToken := strconv.FormatInt(num, 10)
				// replace the hexadecimal or binary string with decimal string
				tokens[indexToProcess] = newPreviousToken
			}
			// remove the (hex) or (bin) markers from the slice
			tokens = append(tokens[:i], tokens[i+1:]...)
			// decrement index since we removed an element
			i--
		}
	}
	return tokens
}

func handleUpLowCap(tokens []string) []string {
	//iterate through all tokens
	for i := 0; i < len(tokens); i++ {
		currentToken := tokens[i]
		//check if current token is lowercase or uppercase or capitalize conversion marker
		if currentToken == "(up)" || currentToken == "(low)" || currentToken == "(cap)" {
			//get the index of the previous token
			indexToProcess := i - 1
			//go back the previous token if its not a word
			for indexToProcess >= 0 && !isWord(tokens[indexToProcess]) {
				indexToProcess--
			}
			//ensure a previous token exists
			if indexToProcess >= 0 {
				previousToken := tokens[indexToProcess]
				var newPreviousToken string
				switch currentToken {
				case "(up)":
					//convert to uppercase
					newPreviousToken = strings.ToUpper(previousToken)
				case "(low)":
					//convert to lower case
					newPreviousToken = strings.ToLower(previousToken)
				case "(cap)":
					//capitalize
					newPreviousToken = Capitalize(previousToken)
				}
				//replace the previous token with the new token
				tokens[indexToProcess] = newPreviousToken
			}
			//remove the (up) or (low) or (cap) markers from the slice
			tokens = append(tokens[:i], tokens[i+1:]...)
			//decrement the index since we removed an element
			i--
		}
	}
	return tokens
}

func handleNumberedUpLowCap(tokens []string) []string {
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		// check for marker start
		if token == "(up," || token == "(low," || token == "(cap," {
			// check if next token is valid
			if i+1 < len(tokens) && strings.HasSuffix(tokens[i+1], ")") {
				rawNumber := strings.TrimSuffix(tokens[i+1], ")")
				count, err := strconv.Atoi(rawNumber)
				// only process if its a valid positive number
				if err == nil {
					// loop over the tokens and convert
					for j := i - 1; count > 0; j-- {
						//ensure index is higher than or equals 0
						if j < 0 {
							break
						}
						//ensure the token is a word
						if !isWord(tokens[j]) {
							continue
						}
						switch token {
						case "(up,":
							tokens[j] = strings.ToUpper(tokens[j])
						case "(low,":
							tokens[j] = strings.ToLower(tokens[j])
						case "(cap,":
							tokens[j] = Capitalize(tokens[j])
						}
						count--
					}
					// remove the tag and the number
					tokens = append(tokens[:i], tokens[i+2:]...)
					// decrement i since we removed two elements
					i--
				}
			}
		}
	}
	return tokens
}
