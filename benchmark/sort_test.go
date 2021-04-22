package main

import (
	"sort"
	"testing"
)

func BenchmarkSortString(b *testing.B) {
	s := []string{"heart", "lungs", "brain", "kidneys", "pancreas"}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sort.Strings(s)
	}

}
