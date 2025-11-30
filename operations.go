package main

import (
	"fmt"
	"go-reloaded/textprocessor"
	"os"
	"strings"
)

// checks command line arguments and returns true
func validateArgs() bool {
	if len(os.Args) == 1 {
		fmt.Println("input and output files names missing")
		return false
	}
	if len(os.Args) == 2 {
		fmt.Println("output file's name missing")
		return false
	}
	if len(os.Args) == 3 {
		return true
	}
	fmt.Println("too many arguments")
	return false
}

// reads the content of a file and returns it as a string
func readFile(filename string) (string, error) {
	info, err := os.Stat(filename)
	if err != nil {
		return "", err
	}
	const maxSize = 100 * 1024
	if info.Size() > maxSize {
		return "", fmt.Errorf("input file is too large")
	}
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// splits text into lines and then into tokens
func tokenize(text string) [][]string {
	lines := strings.Split(text, "\n")
	linesTokens := [][]string{}
	for _, line := range lines {
		tokens := strings.Fields(line)
		linesTokens = append(linesTokens, tokens)
	}
	return linesTokens
}

// applies a processing function to all lines
func applyProcessor(lines [][]string, processor func([]string) []string) [][]string {
	result := [][]string{}
	for _, lineTokens := range lines {
		result = append(result, processor(lineTokens))
	}
	return result
}

// applies all text processing stages in sequence
func processAllStages(linesTokens [][]string) [][]string {
	processed := applyProcessor(linesTokens, textprocessor.HandleHexAndBin)
	processed = applyProcessor(processed, textprocessor.HandleUpLowCap)
	processed = applyProcessor(processed, textprocessor.HandleNumberedUpLowCap)
	processed = applyProcessor(processed, textprocessor.HandlePunctuation)
	processed = applyProcessor(processed, textprocessor.HandleQuotes)
	processed = applyProcessor(processed, textprocessor.HandleVowels)
	return processed
}

// reconstructText joins tokens back into lines and lines into final text
func reconstructText(tokens [][]string) string {
	finalTextLines := []string{}
	for _, lineTokens := range tokens {
		str := strings.Join(lineTokens, " ")
		finalTextLines = append(finalTextLines, str)
	}
	return strings.Join(finalTextLines, "\n")
}

// writes content to a file
func writeFile(filename string, content string) {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Println("failed to write output file")
		return
	}
}
