# Algorithms in Go

This repository contains implementations of various algorithms in Go, organized by category with comprehensive benchmark tests.

## 🔍 Searching Algorithms

The `searching/` directory contains implementations of various search algorithms, each in its own package with benchmark tests.

### Available Algorithms

| Algorithm | Time Complexity | Space Complexity | Best Use Case |
|-----------|----------------|------------------|---------------|
| **Linear Search** | O(n) | O(1) | Unsorted arrays, small datasets |
| **Binary Search** | O(log n) | O(1) | Sorted arrays, most efficient for large datasets |
| **Jump Search** | O(√n) | O(1) | When binary search is costly (linked lists) |
| **Interpolation Search** | O(log log n)* | O(1) | Uniformly distributed sorted arrays |
| **Exponential Search** | O(log n) | O(1) | Unbounded/infinite arrays |
| **Fibonacci Search** | O(log n) | O(1) | When division operations are costly |

*O(n) in worst case for non-uniform distribution

### Package Structure

```
searching/
├── linary/          # Linear Search
├── binary/          # Binary Search  
├── jump/            # Jump Search
├── interpolation/   # Interpolation Search
├── exponential/     # Exponential Search
└── fibonacci/       # Fibonacci Search
```

Each package contains:
- `*.go` - Algorithm implementation
- `*_test.go` - Comprehensive benchmark tests

### Running Benchmarks

```bash
# Navigate to searching directory
cd searching

# Run all search algorithm benchmarks
go test -bench=. ./...

# Run benchmarks with memory allocation stats
go test -bench=. -benchmem ./...

# Run specific algorithm benchmarks
go test -bench=. ./binary
go test -bench=. ./interpolation

# Run benchmarks multiple times for accuracy
go test -bench=. -count=5 ./...
```

### Benchmark Results Example

```
BenchmarkBinarySearchLarge-8             488984020    2.075 ns/op
BenchmarkInterpolationSearchLarge-8      566060540    2.122 ns/op
BenchmarkExponentialSearchLarge-8        47690094     24.53 ns/op
BenchmarkFibonacciSearchLarge-8          42400579     27.93 ns/op
BenchmarkJumpSearchLarge-8               8009985      161.8 ns/op
BenchmarkLinearSearchLarge-8             79647        15573 ns/op
```

### Algorithm Details

#### Linear Search (`linary/`)
- **Description**: Sequentially checks each element until target is found
- **Best for**: Unsorted arrays, small datasets
- **Implementation**: Simple iteration through array

#### Binary Search (`binary/`)
- **Description**: Divides sorted array in half repeatedly
- **Best for**: Large sorted datasets
- **Requirements**: Array must be sorted
- **Implementation**: Iterative approach with left/right pointers

#### Jump Search (`jump/`)
- **Description**: Jumps ahead by fixed steps, then linear search in block
- **Best for**: Large arrays where binary search overhead is costly
- **Step size**: √n (square root of array length)
- **Implementation**: Two-phase search (jumping + linear)

#### Interpolation Search (`interpolation/`)
- **Description**: Estimates position based on value distribution
- **Best for**: Uniformly distributed sorted data
- **Performance**: Can be faster than binary search on uniform data
- **Implementation**: Uses interpolation formula for position estimation

#### Exponential Search (`exponential/`)
- **Description**: Finds range by exponential jumps, then binary search
- **Best for**: Unbounded arrays or when size is unknown
- **Implementation**: Exponential range finding + binary search
- **Also known as**: Doubling search or galloping search

#### Fibonacci Search (`fibonacci/`)
- **Description**: Uses Fibonacci numbers to divide array
- **Best for**: When division operations are expensive
- **Implementation**: Fibonacci sequence for array division
- **Advantage**: Uses only addition and subtraction

### Performance Characteristics

Based on benchmark results with 100,000 elements:

1. **Fastest**: Interpolation Search (~2ns) - for uniform data
2. **Very Fast**: Binary Search (~2ns) - consistently good
3. **Fast**: Exponential Search (~25ns) - good for specific cases
4. **Moderate**: Fibonacci Search (~28ns) - when division is costly
5. **Slower**: Jump Search (~162ns) - scales as √n
6. **Slowest**: Linear Search (~15,573ns) - scales linearly

### Usage Example

```go
package main

import (
    "fmt"
    "searching/binary"
    "searching/linary"
)

func main() {
    arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
    target := 7

    // Binary search (requires sorted array)
    result := binary.Search(arr, target)
    fmt.Printf("Binary search result: %d\n", result) // Output: 3

    // Linear search (works on any array)
    result = linary.Search(arr, target)
    fmt.Printf("Linear search result: %d\n", result) // Output: 3
}
```

## 🚀 Getting Started

1. Clone the repository
2. Navigate to the searching directory: `cd searching`
3. Run tests: `go test -bench=. ./...`
4. Import packages in your Go code as needed

## 📊 Contributing

When adding new search algorithms:
1. Create a new package directory
2. Implement the `Search(arr []int, target int) int` function
3. Add comprehensive benchmark tests using `b.Loop()`
4. Include time/space complexity comments
5. Update this README with algorithm details
