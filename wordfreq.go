package main

import (
	"fmt"
	"regexp"
	"strings"
)

func wordFrequencyCounter(input string) map[string]int {
	frequency := make(map[string]int)

	// Remove punctuation and convert to lowercase
	re := regexp.MustCompile(`[[:punct:]]`)
	cleanedInput := re.ReplaceAllString(input, "")
	cleanedInput = strings.ToLower(cleanedInput)

	// Split the string into words
	words := strings.Fields(cleanedInput)

	// Count the frequency of each word
	for _, word := range words {
		frequency[word]++
	}

	return frequency
}

func main() {
	input := "The function should take in a string as input. The function should return a dictionary with each word in the string as a key and its frequency as the value. Ignore punctuation and case sensitivity. Consider words separated by spaces. Handle contractions and possessives properly."

	result := wordFrequencyCounter(input)

	// Print the result
	for word, count := range result {
		fmt.Printf("%s: %d\n", word, count)
	}
}
