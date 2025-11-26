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

func handlePunctuation(tokens []string) []string {
	//iterate through all tokens
	punctuations := []rune{'.', ',', '!', '?', ':', ';'}
	for i := 0; i < len(tokens); i++ {
		currentToken := tokens[i]
		//iterate through each token
		currentTokenRunes := []rune(currentToken)
		for j := 0; j < len(currentTokenRunes); j++ {
			char := currentTokenRunes[j]
			//check if character is a punctuation
			isPunctuation := false
			for _, punctuation := range punctuations {
				if char == punctuation {
					isPunctuation = true
					break
				}
			}
			if isPunctuation {
				//ensure a previous token exists
				if i-1 >= 0 {
					//add the punctuation to the previous token
					tokens[i-1] = tokens[i-1] + string(char)
				} else {
					//create a new token at the beginning of tokens
					newToken := string(char)
					tokens = append([]string{newToken}, tokens...)
					//increment the index since we added an element
					i++
				}
				//remove the punctuation from the current runes slice and update the current token value
				currentTokenRunes = append(currentTokenRunes[:j], currentTokenRunes[j+1:]...)
				tokens[i] = string(currentTokenRunes)
				//decrement the index since we removed an element
				j--
			} else {
				break
			}
		}
		//remove current token if it becomes empty
		if tokens[i] == "" {
			tokens = append(tokens[:i], tokens[i+1:]...)
			//decrement the index since we removed an element
			i--
		}
	}
	//clean up function
	for i := 0; i < len(tokens); i++ {
		currentTokenRunes := []rune(tokens[i])
		//iterate through the current token\
		for j := 0; j < len(currentTokenRunes); j++ {
			char := currentTokenRunes[j]
			//check if it's a normal character
			if !isPunctuation(char) {
				//ensure a previous token exists and is a punctuation
				if j-1 >= 0 && isPunctuation(currentTokenRunes[j-1]) {
					//split the current token into 2 tokens
					firstToken := currentTokenRunes[:j]
					secondToken := currentTokenRunes[j:]
					//remove current token from tokens and add firstToken and secondToken to tokens
					tempSlice := append([]string{string(firstToken), string(secondToken)}, tokens[i+1:]...)
					tokens = append(tokens[:i], tempSlice...)
					break
				}
			}
		}
	}
	return tokens
}

/* func handleQuotes(tokens []string) []string {
	//join with space
	strRunes := []rune(strings.Join(tokens, " "))
	for i := 0; i < len(strRunes); i++ {
		//if its a quote
		if strRunes[i] == '\'' {
			//ensure it's not in the middle of a word
			if !((i-1 >= 0 && isWordChar(strRunes[i-1])) && (i+1 < len(strRunes) && isWordChar(strRunes[i+1]))) {
				//search for it's next pair
				//if there is no next char break
				if i+1 >= len(strRunes) {
					break
				} else {
					for j := i + 1; j < len(strRunes); j++ {
						//if it's a quote
						if strRunes[j] == '\'' {
							//ensure it's not in the middle of a word
							if !(i-1 >= 0 && isWordChar(strRunes[i-1])) && (i+1 < len(strRunes) && isWordChar(strRunes[i+1])) {
								firstPart := strRunes[:i]
								betweenQuotes := strRunes[i : j+1]
								lastPart := strRunes[j+1:]
								//add a space at the end of first part if there isn't one
								if firstPart[len(firstPart)-1] != ' ' {
									firstPart = append(firstPart, ' ')
								}
								//remove quotes from between quotes
								betweenQuotes = betweenQuotes[1 : len(betweenQuotes)-1]
								//trim space at the start and at the end of between quotes
								betweenQuotes = []rune(strings.TrimPrefix(string(betweenQuotes), " "))
								betweenQuotes = []rune(strings.TrimSuffix(string(betweenQuotes), " "))
								//add a space at the start of the last part if there isn't one
								if lastPart[0] == ' ' {
									lastPart = append([]rune{' '}, lastPart...)
								}
								strRunes = []rune(string(firstPart) + string(betweenQuotes) + string(lastPart))
								i = j + 1
								break
							}
						}
					}
				}
			}
		}
	}
	str := string(strRunes)
	tokens = strings.Fields(str)
	return tokens
} */

func handleQuotes(tokens []string) []string {
	if len(tokens) == 0 {
		return tokens
	}
	
	strRunes := []rune(strings.Join(tokens, " "))
	
	for i := 0; i < len(strRunes); i++ {
		// if it's a quote
		if strRunes[i] == '\'' {
			// ensure it's not in the middle of a word (apostrophe)
			if !((i-1 >= 0 && isWordChar(strRunes[i-1])) && (i+1 < len(strRunes) && isWordChar(strRunes[i+1]))) {
				// search for its closing pair
				if i+1 >= len(strRunes) {
					break
				}
				
				foundPair := false
				for j := i + 1; j < len(strRunes); j++ {
					// if it's a quote
					if strRunes[j] == '\'' {
						// ensure it's not in the middle of a word (apostrophe)
						if !((j-1 >= 0 && isWordChar(strRunes[j-1])) && (j+1 < len(strRunes) && isWordChar(strRunes[j+1]))) {
							// extract the content between quotes (without the quotes themselves)
							betweenQuotes := string(strRunes[i+1 : j])
							
							// trim spaces only at the start and end
							trimmed := strings.TrimSpace(betweenQuotes)
							
							// Count leading and trailing spaces to remove
							leadingSpaces := len(betweenQuotes) - len(strings.TrimLeft(betweenQuotes, " "))
							trailingSpaces := len(betweenQuotes) - len(strings.TrimRight(betweenQuotes, " "))
							
							// Remove the spaces after opening quote
							if leadingSpaces > 0 {
								strRunes = append(strRunes[:i+1], strRunes[i+1+leadingSpaces:]...)
								j -= leadingSpaces
							}
							
							// Remove the spaces before closing quote
							if trailingSpaces > 0 {
								strRunes = append(strRunes[:j-trailingSpaces], strRunes[j-trailingSpaces:]...)
								j -= trailingSpaces
							}
							
							i = j
							foundPair = true
							break
						}
					}
				}
				
				if !foundPair {
					// No closing quote found, move on
					continue
				}
			}
		}
	}
	
	str := string(strRunes)
	tokens = strings.Fields(str)
	return tokens
}