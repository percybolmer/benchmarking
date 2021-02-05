package benching

import (
	"testing"
)

// insertXPreallocIntSlice is used to add X amount of items into a []int
func insertXPreallocIntSlice(x int, b *testing.B) {
	// Initalize a Slice and insert X amount of Items
	testSlice := make([]int, x)
	// reset timer
	b.ResetTimer()
	for i := 0; i < x; i++ {
		testSlice[i] = i
	}
}

// BenchmarkInsertIntSlicePreAllooc1000000 benchmarks the speed of inserting 1000000 integers into the Slice.
func BenchmarkInsertIntSlicePreAllooc1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXPreallocIntSlice(1000000, b)
	}
}

// BenchmarkInsertIntSlicePreAllooc1000000 benchmarks the speed of inserting 100000 integers into the Slice.
func BenchmarkInsertIntSlicePreAllooc100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXPreallocIntSlice(100000, b)
	}
}

// BenchmarkInsertIntSlicePreAllooc10000 benchmarks the speed of inserting 10000 integers into the Slice.
func BenchmarkInsertIntSlicePreAllooc10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXPreallocIntSlice(10000, b)
	}
}

// BenchmarkInsertIntSlicePreAllooc1000 benchmarks the speed of inserting 1000 integers into the Slice.
func BenchmarkInsertIntSlicePreAllooc1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXPreallocIntSlice(1000, b)
	}
}

// BenchmarkInsertIntSlicePreAllooc100 benchmarks the speed of inserting 100 integers into the Slice.
func BenchmarkInsertIntSlicePreAllooc100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXPreallocIntSlice(100, b)
	}
}
