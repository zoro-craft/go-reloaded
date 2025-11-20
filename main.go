package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//---------handling text extraction---------
	//verify that input and output file names are provided
	if len(os.Args) == 1 {
		fmt.Println("input and output files names missing")
	}
	if len(os.Args) == 2 {
		fmt.Println("output file's name missing")
	}
	if len(os.Args) == 3 {
		//collect input and output files names
		iFileName := os.Args[1]
		//oFIleName := os.Args[2]
		//collect text from the input file
		textAsBytes, err := os.ReadFile(iFileName)
		if err != nil {
			fmt.Println("file does not exist")
			return
		}
		//convert bytes to a string
		text := string(textAsBytes)
		//---------splitting text---------
		//split by new lines
		lines := strings.Split(text, "\n")
		//split each line by space
		linesTokens := [][]string{}
		for _, line := range lines {
			lineTokens := strings.Split(line, " ")
			linesTokens = append(linesTokens, lineTokens)
		}
		fmt.Print(linesTokens)
		//---------process text---------
		//handle hex and bin
		
	}
}
