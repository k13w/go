package main

import (
	"testing"
)

func passByReference(x *int) {
	*x++
}

func passByValue(x int) {
	x++
}

func BenchmarkPassByReference(b *testing.B) {
	value := 0
	for i := 0; i < b.N; i++ {
		passByReference(&value)
	}
}

func BenchmarkPassByValue(b *testing.B) {
	value := 0
	for i := 0; i < b.N; i++ {
		passByValue(value)
	}
}

func checkPerformanceMemoryRegister() {
	testing.Benchmark(BenchmarkPassByReference)
	testing.Benchmark(BenchmarkPassByValue)
}
