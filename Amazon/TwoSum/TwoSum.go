package amazon

import (
)

// TwoSum returns the indices of the two numbers such that they add up to a specific target.
func TwoSum(nums []int, target int) []int {
	// Create a map to store the index of each number
	indexMap := make(map[int]int)

	// Iterate through the numbers
	for i, num := range nums {
		// Calculate the complement
		complement := target - num

		// Check if the complement is in the map
		if index, ok := indexMap[complement]; ok {
			// If it is, return the indices
			return []int{index, i}
		}

		// Otherwise, add the number to the map
		indexMap[num] = i
	}

	// If no solution is found, return an empty slice
	return []int{}
}