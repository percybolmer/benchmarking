package benching

import (
	"testing"
)

// selectXIntSlice is used to Select X amount of times from a Slice
func selectXIntSlice(x int, b *testing.B) {
	// Initialize Slice and Insert X amount of items
	testSlice := make([]int, x)
	// Reset timer after Initalizing Slice, that's not what we want to test
	for i := 0; i < x; i++ {
		// Insert value of I into I key.
		testSlice[i] = i
	}
	// holder is a holder that we use to hold the found int, we cannot grab from a Slice without storing the result
	var holder int
	b.ResetTimer()
	for i := 0; i < x; i++ {
		// Select from Slice
		holder = testSlice[i]
	}
	// Compiler wont let us get away with an unused holder, so a quick check will trick the compiler
	if holder != 0 {

	}
}

// BenchmarkSelectIntSlice1000000 benchmarks the speed of selecting 1000000 items from the Slice.
func BenchmarkSelectIntSlice1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXIntSlice(1000000, b)
	}
}

// BenchmarkSelectIntSlice100000 benchmarks the speed of selecting 100000 items from the Slice.
func BenchmarkSelectIntSlice100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXIntSlice(100000, b)
	}
}

// BenchmarkSelectIntSlice10000 benchmarks the speed of selecting 10000 items from the Slice.
func BenchmarkSelectIntSlice10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXIntSlice(10000, b)
	}
}

// BenchmarkSelectIntSlice1000 benchmarks the speed of selecting 1000 items from the Slice.
func BenchmarkSelectIntSlice1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXIntSlice(1000, b)
	}
}

// BenchmarkSelectIntSlice100 benchmarks the speed of selecting 100 items from the Slice.
func BenchmarkSelectIntSlice100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXIntSlice(100, b)
	}
}
