package benching

import (
	"testing"
)

// selectXIntMap is used to Select X amount of times from a map
func selectXIntMap(x int, b *testing.B) {
	// Initialize Map and Insert X amount of items
	testmap := make(map[int]int, x)
	// Reset timer after Initalizing map, that's not what we want to test
	for i := 0; i < x; i++ {
		// Insert value of I into I key.
		testmap[i] = i
	}
	// holder is a holder that we use to hold the found int, we cannot grab from a map without storing the result
	var holder int
	b.ResetTimer()
	for i := 0; i < x; i++ {
		// Select from map
		holder = testmap[i]
	}
	// Compiler wont let us get away with an unused holder, so a quick check will trick the compiler
	if holder != 0 {

	}
}

// BenchmarkSelectIntMap1000000 benchmarks the speed of selecting 1000000 items from the map.
func BenchmarkSelectIntMap1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXIntMap(1000000, b)
	}
}

// BenchmarkSelectIntMap100000 benchmarks the speed of selecting 100000 items from the map.
func BenchmarkSelectIntMap100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXIntMap(100000, b)
	}
}

// BenchmarkSelectIntMap10000 benchmarks the speed of selecting 10000 items from the map.
func BenchmarkSelectIntMap10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXIntMap(10000, b)
	}
}

// BenchmarkSelectIntMap1000 benchmarks the speed of selecting 1000 items from the map.
func BenchmarkSelectIntMap1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXIntMap(1000, b)
	}
}

// BenchmarkSelectIntMap100 benchmarks the speed of selecting 100 items from the map.
func BenchmarkSelectIntMap100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXIntMap(100, b)
	}
}
