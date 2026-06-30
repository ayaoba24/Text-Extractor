package main

import (
	"fmt"
	"strings"
	"unicode"
)

func textStats(text string) {
	words := len(strings.Fields(text))
	characters := len([]rune(text)) 

	letters := 0
	punctuation := 0
	for _, c := range text {
		if unicode.IsLetter(c) {
			letters++
		} else if unicode.IsPunct(c) {
			punctuation++
		}
	}

	lines := strings.Count(text, "\n") + 1

	fmt.Println("Words:", words)
	fmt.Println("Characters:", characters)
	fmt.Println("Letters:", letters)
	fmt.Println("Punctuation:", punctuation)
	fmt.Println("Lines:", lines)
}

