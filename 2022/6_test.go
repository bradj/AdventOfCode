package twenttwo

import (
	"testing"

	"github.com/bradj/AdventOfCode/util"
)

func BenchmarkD6p2(b *testing.B) {
	items := util.GetItems("6.input")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		D6p2(items)
	}
}

func BenchmarkD6p2_2(b *testing.B) {
	items := util.GetItems("6.input")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		D6p2_2(items)
	}
}
