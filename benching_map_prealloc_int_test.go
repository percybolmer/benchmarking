package benching

import (
	"testing"
)

// insertXPreallocIntMap is used to add X amount of items into a Map[int]int
func insertXPreallocIntMap(x int, b *testing.B) {
	// Initialize Map and Insert X amount of items and Prealloc the size to X
	testmap := make(map[int]int, x)
	// Reset timer after Initalizing map, that's not what we want to test
	b.ResetTimer()
	for i := 0; i < x; i++ {
		// Insert value of I into I key.
		testmap[i] = i
	}
}

// BenchmarkInsertIntMapPreAlloc1000000 benchmarks the speed of inserting 1000000 integers into the map.
func BenchmarkInsertIntMapPreAlloc1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXPreallocIntMap(1000000, b)
	}
}

// BenchmarkInsertIntMapPreAlloc100000 benchmarks the speed of inserting 100000 integers into the map.
func BenchmarkInsertIntMapPreAlloc100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXPreallocIntMap(100000, b)
	}
}

// BenchmarkInsertIntMapPreAlloc10000 benchmarks the speed of inserting 10000 integers into the map.
func BenchmarkInsertIntMapPreAlloc10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXPreallocIntMap(10000, b)
	}
}

// BenchmarkInsertIntMapPreAlloc1000 benchmarks the speed of inserting 1000 integers into the map.
func BenchmarkInsertIntMapPreAlloc1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXPreallocIntMap(1000, b)
	}
}

// BenchmarkInsertIntMapPreAlloc100 benchmarks the speed of inserting 100 integers into the map.
func BenchmarkInsertIntMapPreAlloc100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXPreallocIntMap(100, b)
	}
}
