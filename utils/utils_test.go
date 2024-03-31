package utils

import "testing"

func BenchmarkGetARandom(b *testing.B) {

	for i := 0; i < b.N; i++ {
		GetARandom(4)
	}
}
