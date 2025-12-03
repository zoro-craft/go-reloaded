package main

import (
	"slices"
	"testing"
)

func TestTokenize(t *testing.T) {
	inputFile := "sample.txt"
	text, err := readFile(inputFile)
	if err != nil {
		t.Skipf("error reading the %s file", inputFile)
	}
	got := tokenize(text)
	gotJoined := reconstructText(got)
	want := text
	if gotJoined != want {
		t.Error("error tokenizing")
	}
}

func TestProcessAllStages(t *testing.T) {
	inputFile := "sample.txt"
	expectedOutputFile := "expectedResult.txt"
	input, _ := readFile(inputFile)
	expectedOutput, _ := readFile(expectedOutputFile)
	linesTokens := tokenize(input)
	want := tokenize(expectedOutput)
	got := processAllStages(linesTokens)
	inner := func(a, b []string) bool {
		return slices.Equal(a, b)
	}
	if !slices.EqualFunc(got, want, inner) {
		t.Error()
	}
}
