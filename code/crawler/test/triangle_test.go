package test

import (
	"testing"
)

//// 必须Test开头
//func TestTriangle(t *testing.T) {
//	tests := []struct{ a, b, c int }{
//		{3, 4, 5},
//		{5, 12, 13},
//		{8, 15, 17},
//		{12, 33, 4},
//		{20000, 30000, 50000},
//	}
//
//	for _, te := range tests {
//		if actual := Triangle(te.a, te.b); actual != te.c {
//			t.Errorf("triangle(%d,%d) got %d expected %d", te.a, te.b, actual, te.c)
//		}
//	}
//
//}

// 性能测试，必须Benchmark开头
func BenchmarkTriangle(b *testing.B) {
	// 执行多少由Benchmark决定
	for i := 0; i < b.N; i++ {
		if actual := Triangle(3, 4); actual != 5 {
			b.Errorf("triangle(%d,%d) got %d expected %d", 3, 4, actual, 5)
		}
	}

}
