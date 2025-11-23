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
			linesTokensHexBinHandled = append(linesTokensHexBinHandled, handleHexAndBin(lineTokens))
		}
		//fmt.Println(linesTokensHexBinHandled)
		// process the simple uppercase and lowercase and capitalize conversions
		linesUpLowCapHandled := [][]string{}
		for _, lineTokens := range linesTokensHexBinHandled {
			linesUpLowCapHandled = append(linesUpLowCapHandled, handleUpLowCap(lineTokens))
		}
		//fmt.Println(linesUpLowCapHandled)
		// process the numbered uppercase and lowercase and capitalize conversions
		linesNumberedUpLowCapHandled := [][]string{}
		for _, lineTokens := range linesUpLowCapHandled {
			linesNumberedUpLowCapHandled = append(linesNumberedUpLowCapHandled, handleNumberedUpLowCap(lineTokens))
		}
		//process the punctuations
		handledPunctuation := [][]string{}
		for _, lineTokens := range linesNumberedUpLowCapHandled {
			handledPunctuation = append(handledPunctuation, handlePunctuation(lineTokens))
		}
		//current process
		finalTextLines := []string{}
		for _, token := range handledPunctuation {
			str := strings.Join(token, " ")
			finalTextLines = append(finalTextLines, str)
		}
		finalText := strings.Join(finalTextLines, "\n")
		fmt.Println(finalText)
		os.WriteFile(outputFile, []byte(finalText), 0644)
	}
}
