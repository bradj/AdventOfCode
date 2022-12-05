package twenttwo

import (
	"testing"

	"github.com/bradj/AdventOfCode/util"
)

func BenchmarkD3p1(b *testing.B) {
	items := util.GetItems("3.input")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		D3p1(items)
	}
}
