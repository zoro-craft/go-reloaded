package main

import (
	"os"
)

func main() {
	isValid := validateArgs()
	if !isValid {
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	text, err := readFile(inputFile)
	if err != nil {
		return
	}
	linesTokens := tokenize(text)
	processedTokens := processAllStages(linesTokens)
	finalText := reconstructText(processedTokens)
	writeFile(outputFile, finalText)
}
