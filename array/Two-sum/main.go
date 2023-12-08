package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// Create a map to store the complement of each number and its index
	numMap := make(map[int]int)

	for i, num := range nums {
		/*
			The line complement := target - num calculates the value that, when added to the current num, will result in the target sum. If this complement value is found in the array later on, it means that adding the current num to that complement will result in the target sum.

			Here's a breakdown of the logic:

			If num + complement == target, then complement == target - num.
			So, by keeping track of the complements in a map as you iterate through the array, you can efficiently find pairs of numbers whose sum equals the target. If you find a number whose complement is already in the map, you have found a pair that satisfies the condition, and you can return the indices of those two numbers.
		*/
		complement := target - num

		// Check if the complement exists in the map
		if index, found := numMap[complement]; found {
			fmt.Println(index)
			// If found, print the indices of the two numbers
			return []int{index, i}
		}

		// Otherwise, store the current number and its index in the map
		numMap[num] = i
	}

	// If no solution is found, return an empty slice or handle it accordingly
	return nil
}

func main() {
	// Example usage:
	nums1 := []int{2, 5, 7, 11, 15}
	target1 := 9
	twoSumA := twoSum(nums1, target1)
	fmt.Println(twoSumA)

	nums2 := []int{3, 2, 4}
	target2 := 6
	twoSumB := twoSum(nums2, target2)
	fmt.Println(twoSumB)

	nums3 := []int{3, 3}
	target3 := 6
	twoSumC := twoSum(nums3, target3)
	fmt.Println(twoSumC)

}
