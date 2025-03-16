package main

import "testing"

func BenchmarkCollatzOddTraditional(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = 3*i + 1
	}
}

func BenchmarkCollatzEvenTraditional(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = i / 2
	}
}

func BenchmarkCollatzOddBitwise(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ((i << 1) | 1) + i
	}
}

func BenchmarkCollatzEvenBitwise(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = i >> 1
	}
}
