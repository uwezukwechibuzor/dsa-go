package main

import (
	"fmt"
	"math"
)

func MaxSubArraySum(data []int) int {
	size := len(data)
	maxSoFar := 0
	maxEndingHere := 0
	for i := 0; i < size; i++ {
		//adding the values till current value
		maxEndingHere = maxEndingHere + data[i]
		// reset index if its negative
		if maxEndingHere < 0 {
			maxEndingHere = 0
		}

		//reset maximum global if its smaller then current one
		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}
	}
	return maxSoFar
}

func MinSubArraySum(data []int) int {
	size := len(data)
	minSoFar := math.MaxInt32
	minEndingHere := 0
	for i := 0; i < size; i++ {
		// adding the values till the current value
		minEndingHere = minEndingHere + data[i]
		// reset index if it's positive
		if minEndingHere > 0 {
			minEndingHere = 0
		}

		// reset minimum global if it's greater than the current one
		if minSoFar > minEndingHere {
			minSoFar = minEndingHere
		}
	}
	return minSoFar
}

// Testing code
func main() {
	data := []int{1, -2, 3, 4, -4, 6, -14, 6, 2}
	fmt.Println("Max sub array sum :", MaxSubArraySum(data))
	fmt.Println("Min sub array sum :", MinSubArraySum(data))
}
