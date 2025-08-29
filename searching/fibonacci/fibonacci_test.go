package fibonacci

import "testing"

func BenchmarkFibonacciSearchSmall(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7
	b.ResetTimer()

	for b.Loop() {
		Search(arr, target)
	}
}

func BenchmarkFibonacciSearchMedium(b *testing.B) {
	arr := make([]int, 1000)

	for i := range arr {
		arr[i] = i + 1
	}

	target := 500
	b.ResetTimer()

	for b.Loop() {
		Search(arr, target)
	}
}

func BenchmarkFibonacciSearchLarge(b *testing.B) {
	arr := make([]int, 100000)

	for i := range arr {
		arr[i] = i + 1
	}

	target := 50000
	b.ResetTimer()

	for b.Loop() {
		Search(arr, target)
	}
}

func BenchmarkFibonacciSearchWorstCase(b *testing.B) {
	arr := make([]int, 10000)

	for i := range arr {
		arr[i] = i + 1
	}

	target := -1
	b.ResetTimer()

	for b.Loop() {
		Search(arr, target)
	}
}
