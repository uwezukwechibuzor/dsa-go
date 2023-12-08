package main

import (
	"fmt"
)

func plusMinus(arr []int32) {
	n := float64(len(arr))
	positive, negative, zero := 0, 0, 0

	for _, num := range arr {
		if num > 0 {
			positive++
		} else if num < 0 {
			negative++
		} else {
			zero++
		}
	}

	// Calculate the ratios
	ratioPositive := float64(positive) / n
	ratioNegative := float64(negative) / n
	ratioZero := float64(zero) / n

	// Print the ratios with 6 decimal places
	fmt.Printf("%.6f\n", ratioPositive)
	fmt.Printf("%.6f\n", ratioNegative)
	fmt.Printf("%.6f\n", ratioZero)
}

func main() {
	// Example usage
	arr := []int32{-4, 3, -9, 0, 4, 1, 1, 0}
	plusMinus(arr)
}
