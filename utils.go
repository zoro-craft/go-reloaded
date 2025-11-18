package main

import "strings"

func uniteFlags(strSlice []string) (string, []string) {
	//make flags
	unitedFlagsSlice := []string{}
	flag := ""
	isError:= false
	for _, element := range strSlice {
		if strings.HasPrefix(element, "(") && strings.HasSuffix(element, ")") {
			unitedFlagsSlice = append(unitedFlagsSlice, element)
		} else if !strings.HasPrefix(element, "(") && !strings.HasSuffix(element, ")") {
			unitedFlagsSlice = append(unitedFlagsSlice, element)
		} else if strings.HasPrefix(element, "(") && !strings.HasSuffix(element, ")") {
			flag += element
		} else if strings.HasSuffix(element, ")") && !strings.HasPrefix(element, "(") {
			unitedFlagsSlice = append(unitedFlagsSlice, flag)
			flag = ""
		}
	}

}
