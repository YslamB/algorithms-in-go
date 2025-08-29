package linary

// Search performs linear search on a sorted array
// Returns the index of the target element, or -1 if not found
// Time complexity: O(n) - linear time in the worst case, where n is the length of the array.
// Space complexity: O(1) - constant space, as no additional space is used.
func Search(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}
