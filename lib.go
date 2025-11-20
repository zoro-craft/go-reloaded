package main

import "strconv"

// handle (hex) and (bin)
func handleHexAndBin(lineTokens []string) []string {
	//iterate over the lineToken
	for i := 0; i < len(lineTokens); i++ {
		lineToken := lineTokens[i]
		//check if lineToken equals (hex)
		if lineToken == "(hex)" {
			//transform previous lineToken if it exists
			previousIndex := i - 1
			if previousIndex >= 0 {
				previousLineToken := lineTokens[previousIndex]
				num, err := strconv.ParseInt(previousLineToken, 16, 64)
				if err != nil {
					continue
				}
			}
		}
	}

	return nil
}
