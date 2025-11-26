package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// --------- validate command line arguments ---------

	// ensure both input and output file names are provided
	if len(os.Args) == 1 {
		fmt.Println("input and output files names missing")
		return
	}
	if len(os.Args) == 2 {
		fmt.Println("output file's name missing")
		return
	}
	if len(os.Args) == 3 {
		// extract input and output file names from arguments
		inputFile := os.Args[1]
		outputFile := os.Args[2]

		// --------- read input file ---------
		// read the entire file content as bytes
		content, err := os.ReadFile(inputFile)
		if err != nil {
			fmt.Println("file does not exist")
			return
		}
		// convert file content from bytes to string
		text := string(content)
		// --------- tokenize text ---------
		// split content into individual lines
		lines := strings.Split(text, "\n")
		// split each line into words (tokens)
		linesTokens := [][]string{}
		for _, line := range lines {
			tokens := strings.Fields(line)
			linesTokens = append(linesTokens, tokens)
		}
		// --------- process tokens ---------
		// process hexadecimal and binary conversions
		linesTokensHexBinHandled := [][]string{}
		for _, lineTokens := range linesTokens {
			copyLine := append([]string{}, lineTokens...)
			linesTokensHexBinHandled = append(linesTokensHexBinHandled, handleHexAndBin(copyLine))
		}
		fmt.Println("linesUpLowCapHandled", linesTokensHexBinHandled)
		// process the simple uppercase and lowercase and capitalize conversions
		linesUpLowCapHandled := [][]string{}
		for _, lineTokens := range linesTokensHexBinHandled {
			copyLine := append([]string{}, lineTokens...)
			linesUpLowCapHandled = append(linesUpLowCapHandled, handleUpLowCap(copyLine))
		}
		fmt.Println("linesUpLowCapHandled", linesUpLowCapHandled)
		// process the numbered uppercase and lowercase and capitalize conversions
		linesNumberedUpLowCapHandled := [][]string{}
		for _, lineTokens := range linesUpLowCapHandled {
			copyLine := append([]string{}, lineTokens...)
			linesNumberedUpLowCapHandled = append(linesNumberedUpLowCapHandled, handleNumberedUpLowCap(copyLine))
		}
		fmt.Println("linesNumberedUpLowCapHandled", linesNumberedUpLowCapHandled)
		//process the punctuations
		handledPunctuation := [][]string{}
		for _, lineTokens := range linesNumberedUpLowCapHandled {
			copyLine := append([]string{}, lineTokens...)
			handledPunctuation = append(handledPunctuation, handlePunctuation(copyLine))
		}
		fmt.Println("handledPunctuation", handledPunctuation)

		//process the quotes
		handledQuotes := [][]string{}
		for _, lineTokens := range handledPunctuation {
			copyLine := append([]string{}, lineTokens...)
			handledQuotes = append(handledQuotes, handleQuotes(copyLine))
		}
		//current process
		finalTextLines := []string{}
		for _, token := range handledQuotes {
			str := strings.Join(token, " ")
			finalTextLines = append(finalTextLines, str)
		}
		finalText := strings.Join(finalTextLines, "\n")
		//fmt.Println(finalText)
		os.WriteFile(outputFile, []byte(finalText), 0644)
	}
}
