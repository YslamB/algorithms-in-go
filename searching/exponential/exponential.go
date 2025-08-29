package exponential

// Search performs exponential search on a sorted array
// Returns the index of the target element, or -1 if not found
// Time complexity: O(log n) - logarithmic time in the worst case, where n is the length of the array.
// Space complexity: O(1) - constant space, as no additional space is used.
func Search(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	if arr[0] == target {
		return 0
	}

	i := 1
	for i < n && arr[i] <= target {
		i = i * 2
	}

	return binarySearch(arr, i/2, min(i, n-1), target)
}

func binarySearch(arr []int, left, right, target int) int {
	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
