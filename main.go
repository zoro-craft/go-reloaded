package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	isValid := validateArgs()
	if !isValid {
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	//validate file extension
	if !strings.HasSuffix(inputFile, ".txt") || !strings.HasSuffix(outputFile, ".txt") {
		fmt.Println("both files must have .txt at the end")
		return
	}
	//validate file name
	if len(inputFile) <= 4 || len(outputFile) <= 4 {
		fmt.Println("both files must have a name")
		return
	}
	text, err := readFile(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	linesTokens := tokenize(text)
	processedTokens := processAllStages(linesTokens)
	finalText := reconstructText(processedTokens)

	writeFile(outputFile, finalText)
}
