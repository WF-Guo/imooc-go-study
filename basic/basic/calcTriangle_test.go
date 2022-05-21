package basic

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 0},
		{4, 5, 5},
		{30000, 40000, 50000},
	}
	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d, %d); "+
				"got %d; excepted %d", tt.a, tt.b, actual, tt.c)
		}
	}
}

func BenchmarkTest(B *testing.B) {
	a, b, c := 30000, 40000, 50000

	for i := 0; i < B.N; i++ {
		if actual := calcTriangle(a, b); actual != c {
			B.Errorf("calcTriangle(%d, %d); "+
				"got %d; excepted %d", a, b, actual, c)
		}
	}
}
