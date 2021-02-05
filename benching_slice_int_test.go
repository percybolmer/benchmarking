package benching

import (
	"testing"
)

// insertXIntSlice is used to add X amount of items into a []int
func insertXIntSlice(x int, b *testing.B) {
	// Initalize a Slice and insert X amount of Items
	testSlice := make([]int, 0)
	// reset timer
	b.ResetTimer()
	for i := 0; i < x; i++ {
		testSlice = append(testSlice, i)
	}
}

// BenchmarkInsertIntSlice1000000 benchmarks the speed of inserting 1000000 integers into the Slice.
func BenchmarkInsertIntSlice1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXIntSlice(1000000, b)
	}
}

// BenchmarkInsertIntSlice100000 benchmarks the speed of inserting 100000 integers into the Slice.
func BenchmarkInsertIntSlice100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXIntSlice(100000, b)
	}
}

// BenchmarkInsertIntSlice10000 benchmarks the speed of inserting 10000 integers into the Slice.
func BenchmarkInsertIntSlice10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXIntSlice(10000, b)
	}
}

// BenchmarkInsertIntSlice1000 benchmarks the speed of inserting 1000 integers into the Slice.
func BenchmarkInsertIntSlice1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXIntSlice(1000, b)
	}
}

// BenchmarkInsertIntSlice100 benchmarks the speed of inserting 100 integers into the Slice.
func BenchmarkInsertIntSlice100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXIntSlice(100, b)
	}
}
