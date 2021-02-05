package benching

import (
	"testing"
)

// insertXInterfaceMap is used to add X amount of items into a Map[interface]int
func insertXInterfaceMap(x int, b *testing.B) {
	// Initialize Map and Insert X amount of items
	testmap := make(map[interface{}]int, 0)
	// Reset timer after Initalizing map, that's not what we want to test
	b.ResetTimer()
	for i := 0; i < x; i++ {
		// Insert value of I into I key.
		testmap[i] = i
	}
}

// BenchmarkInsertInterfaceMap1000000 benchmarks the speed of inserting 1000000 integers into the map.
func BenchmarkInsertInterfaceMap1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXInterfaceMap(1000000, b)
	}
}

// BenchmarkInsertInterfaceMap100000 benchmarks the speed of inserting 100000 integers into the map.
func BenchmarkInsertInterfaceMap100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXInterfaceMap(100000, b)
	}
}

// BenchmarkInsertInterfaceMap10000 benchmarks the speed of inserting 10000 integers into the map.
func BenchmarkInsertInterfaceMap10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXInterfaceMap(10000, b)
	}
}

// BenchmarkInsertInterfaceMap1000 benchmarks the speed of inserting 1000 integers into the map.
func BenchmarkInsertInterfaceMap1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXInterfaceMap(1000, b)
	}
}

// BenchmarkInsertInterfaceMap100 benchmarks the speed of inserting 100 integers into the map.
func BenchmarkInsertInterfaceMap100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXInterfaceMap(100, b)
	}
}
