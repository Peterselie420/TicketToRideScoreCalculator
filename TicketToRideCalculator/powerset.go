package main


// powerset generates the powerset of a given slice of integers
func powerset(nums []int) [][]int {
	var result [][]int
	generatePowerSet(nums, 0, []int{}, &result)
	return result
}

// generatePowerSet is a helper function for generating the powerset recursively
func generatePowerSet(nums []int, index int, currentSubset []int, result *[][]int) {
	// Add the current subset to the result
	*result = append(*result, append([]int{}, currentSubset...))

	// Iterate over the remaining elements in the slice
	for i := index; i < len(nums); i++ {
		// Include the current element in the current subset
		currentSubset = append(currentSubset, nums[i])
		// Recursively generate subsets starting from the next index
		generatePowerSet(nums, i+1, currentSubset, result)
		// Backtrack and remove the current element from the subset
		currentSubset = currentSubset[:len(currentSubset)-1]
	}
}
