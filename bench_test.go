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
	datas := make([][]*point, b.N)
	for i := 0; i < b.N; i++ {
		cpdata := make([]*point, 10000)
		copy(cpdata, pdata)
		datas[i] = cpdata
	}

	b.ResetTimer()
	var sorter *pointSorter1
	for i := 0; i < b.N; i++ {
		sorter = &pointSorter1{points: datas[i], sc: sc}
		sorter.lessCount = 0
		sort.Sort(sorter)
	}
}

func BenchmarkArraySort(b *testing.B) {
	sc := new(stmtctx.StatementContext)
	datas := make([][]point, b.N)
	for i := 0; i < b.N; i++ {
		cdata := make([]point, 10000)
		copy(cdata, data)
		datas[i] = cdata
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sorter := &pointSorter2{points: datas[i], sc: sc}
		sort.Sort(sorter)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	sc := new(stmtctx.StatementContext)
	datas := make([][]*point, b.N)
	for i := 0; i < b.N; i++ {
		cpdata := make([]*point, 10000)
		copy(cpdata, pdata)
		datas[i] = cpdata
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := new(PointHeap)
		h.sc = sc
		h.lessCount = 0
		heapSort(h, datas[i])
	}
}
