package fibonacci

// Search performs fibonacci search on a sorted array
// Returns the index of the target element, or -1 if not found
// Time complexity: O(log n) - logarithmic time in the worst case, where n is the length of the array.
// Space complexity: O(1) - constant space, as no additional space is used.
func Search(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	fibM2 := 0
	fibM1 := 1
	fibM := fibM2 + fibM1

	for fibM < n {
		fibM2 = fibM1
		fibM1 = fibM
		fibM = fibM2 + fibM1
	}

	offset := -1

	for fibM > 1 {
		i := min(offset+fibM2, n-1)

		if arr[i] < target {
			fibM = fibM1
			fibM1 = fibM2
			fibM2 = fibM - fibM1
			offset = i
		} else if arr[i] > target {
			fibM = fibM2
			fibM1 = fibM1 - fibM2
			fibM2 = fibM - fibM1
		} else {
			return i
		}
	}

	if fibM1 == 1 && offset+1 < n && arr[offset+1] == target {
		return offset + 1
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
