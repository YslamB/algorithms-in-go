package interpolation

// Search performs interpolation search on a sorted array with uniformly distributed values
// Returns the index of the target element, or -1 if not found
// Time complexity: O(log log n) - logarithmic time in the worst case, where n is the length of the array.
// Space complexity: O(1) - constant space, as no additional space is used.
func Search(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high && target >= arr[low] && target <= arr[high] {
		if low == high {
			if arr[low] == target {
				return low
			}
			return -1
		}

		pos := low + ((target-arr[low])*(high-low))/(arr[high]-arr[low])

		if arr[pos] == target {
			return pos
		}

		if arr[pos] < target {
			low = pos + 1
		} else {
			high = pos - 1
		}
	}

	return -1
}
