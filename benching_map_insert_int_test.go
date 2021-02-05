package benching

import (
	"testing"
)

// insertXIntMap is used to add X amount of items into a Map[int]int
func insertXIntMap(x int, b *testing.B) {
	// Initialize Map and Insert X amount of items
	testmap := make(map[int]int, 0)
	// Reset timer after Initalizing map, that's not what we want to test
	b.ResetTimer()
	for i := 0; i < x; i++ {
		// Insert value of I into I key.
		testmap[i] = i
	}
}

// BenchmarkInsertIntMap1000000 benchmarks the speed of inserting 1000000 integers into the map.
func BenchmarkInsertIntMap1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXIntMap(1000000, b)
	}
}

// BenchmarkInsertIntMap100000 benchmarks the speed of inserting 100000 integers into the map.
func BenchmarkInsertIntMap100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXIntMap(100000, b)
	}
}

// BenchmarkInsertIntMap10000 benchmarks the speed of inserting 10000 integers into the map.
func BenchmarkInsertIntMap10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXIntMap(10000, b)
	}
}

// BenchmarkInsertIntMap1000 benchmarks the speed of inserting 1000 integers into the map.
func BenchmarkInsertIntMap1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXIntMap(1000, b)
	}
}

// BenchmarkInsertIntMap100 benchmarks the speed of inserting 100 integers into the map.
func BenchmarkInsertIntMap100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXIntMap(100, b)
	}
}
