# Algorithms in Go

This repository contains implementations of various algorithms in Go, organized by category with comprehensive benchmark tests.

## 🔍 Searching Algorithms

The `searching/` directory contains implementations of various search algorithms, each in its own package with benchmark tests.
Each algorithm is optimized for specific use cases and data distributions, with detailed complexity analysis and performance benchmarks. The implementations prioritize clarity and efficiency, making them suitable for both educational purposes and production use.



### Package Structure

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
