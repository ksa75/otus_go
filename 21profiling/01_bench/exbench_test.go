package exbench

import "testing"

func BenchmarkFast(b *testing.B) {
	// heavy logic

	//b.ResetTimer()

	// b.N = 10 <- ломает бенчмарк
	for i := 0; i < b.N; i++ {
		Fast()
	}
}

func BenchmarkSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Slow()
	}
}
