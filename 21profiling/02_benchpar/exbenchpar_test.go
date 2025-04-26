package benchpar

import "testing"

//func BenchmarkFast(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Fast()
//	}
//}
//
//func BenchmarkSlow(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Slow()
//	}
//}

func BenchmarkParallelFast(b *testing.B) {
	b.SetParallelism(1000)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Fast()
		}
	})
}

func BenchmarkParallelSlow(b *testing.B) {
	b.SetParallelism(1000)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Slow()
		}
	})
}

//BenchmarkParallelFast
//BenchmarkParallelFast-10    	25339687	        46.23 ns/op	       0 B/op	       0 allocs/op
//BenchmarkParallelSlow
//BenchmarkParallelSlow-10    	 2343495	       495.6 ns/op	       0 B/op	       0 allocs/op

//BenchmarkParallelFast
//BenchmarkParallelFast-10    	27530486	        44.12 ns/op	       0 B/op	       0 allocs/op
//BenchmarkParallelSlow
//BenchmarkParallelSlow-10    	 1934572	       619.8 ns/op	       0 B/op	       0 allocs/op

//BenchmarkParallelFast
//BenchmarkParallelFast-10    	27124492	        47.37 ns/op	       0 B/op	       0 allocs/op
//BenchmarkParallelSlow
//BenchmarkParallelSlow-10    	 1903562	       631.0 ns/op	       0 B/op	       0 allocs/op
