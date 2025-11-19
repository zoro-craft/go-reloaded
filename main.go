package main

import (
	"fmt"
	"os"
)

func main() {
	//verify that input and output file names are provided
	if len(os.Args) == 1 {
		fmt.Println("files names missing")
	}
	if len(os.Args) == 2 {
		fmt.Println("output file name missing")
	}
	if len(os.Args) == 3 {
		//collect input and output files names
		iFileName := os.Args[1]
		//oFIleName := os.Args[2]
		//collect text from the input file
		textAsBytes, err := os.ReadFile(iFileName)
		if err != nil {
			fmt.Println("file does not exist")
		} else {
			text := string(textAsBytes)
			fmt.Print(text)
		}
	}

}
