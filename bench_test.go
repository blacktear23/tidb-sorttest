package main

import (
	// "fmt"
	"sort"
	"testing"

	"github.com/pingcap/tidb/sessionctx/stmtctx"
)

var (
	data  []point
	pdata []*point
)

func init() {
	pdata, data = prepareArrays(10000)
}

func BenchmarkPointerArraySort(b *testing.B) {
	sc := new(stmtctx.StatementContext)

	b.ResetTimer()
	var sorter *pointSorter1
	sorter = &pointSorter1{points: pdata, sc: sc}
	for i := 0; i < b.N; i++ {
		sorter.lessCount = 0
		sort.Sort(sorter)
	}
}

func BenchmarkArraySort(b *testing.B) {
	sc := new(stmtctx.StatementContext)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sorter := &pointSorter2{points: data, sc: sc}
		sort.Sort(sorter)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	sc := new(stmtctx.StatementContext)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := new(PointHeap)
		h.sc = sc
		h.lessCount = 0
		heapSort(h, pdata)
	}
}
