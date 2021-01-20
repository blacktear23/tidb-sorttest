package main

import (
	// "fmt"
	"sort"
	"testing"

	"github.com/pingcap/tidb/sessionctx/stmtctx"
)

var (
	data1, data2, data3, data4             []point
	pdata1, pdata2, pdata3, pdata4, pdata5 []*point
)

func init() {
	pdata1, data1 = prepareStringArrays(10000, "utf8mb4_general_ci")
	pdata2, data2 = prepareStringArrays(10000, "utf8mb4_bin")
	pdata3, data3 = prepareIntArrays(10000)
	pdata4, data4 = prepareDecimalArrays(10000)
	pdata5 = prepareStringArrayReverse(10000, "utf8mb4_general_ci")
}

func BenchmarkPointerArraySort_StringUtf8bin(b *testing.B) {
	sc := new(stmtctx.StatementContext)
	datas := make([][]*point, b.N)
	for i := 0; i < b.N; i++ {
		cpdata := make([]*point, 10000)
		copy(cpdata, pdata2)
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

func BenchmarkArraySort_StringUtf8bin(b *testing.B) {
	sc := new(stmtctx.StatementContext)
	datas := make([][]point, b.N)
	for i := 0; i < b.N; i++ {
		cdata := make([]point, 10000)
		copy(cdata, data2)
		datas[i] = cdata
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sorter := &pointSorter2{points: datas[i], sc: sc}
		sort.Sort(sorter)
	}
}

func BenchmarkHeapSort_StringUtf8bin(b *testing.B) {
	sc := new(stmtctx.StatementContext)
	datas := make([][]*point, b.N)
	for i := 0; i < b.N; i++ {
		cpdata := make([]*point, 10000)
		copy(cpdata, pdata2)
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

func BenchmarkPointerArraySort_StringUtf8Mb4(b *testing.B) {
	sc := new(stmtctx.StatementContext)
	datas := make([][]*point, b.N)
	for i := 0; i < b.N; i++ {
		cpdata := make([]*point, 10000)
		copy(cpdata, pdata1)
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

func BenchmarkArraySort_StringUtf8Mb4(b *testing.B) {
	sc := new(stmtctx.StatementContext)
	datas := make([][]point, b.N)
	for i := 0; i < b.N; i++ {
		cdata := make([]point, 10000)
		copy(cdata, data1)
		datas[i] = cdata
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sorter := &pointSorter2{points: datas[i], sc: sc}
		sort.Sort(sorter)
	}
}

func BenchmarkHeapSort_StringUtf8Mb4(b *testing.B) {
	sc := new(stmtctx.StatementContext)
	datas := make([][]*point, b.N)
	for i := 0; i < b.N; i++ {
		cpdata := make([]*point, 10000)
		copy(cpdata, pdata1)
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

func BenchmarkPointerArraySort_Int(b *testing.B) {
	sc := new(stmtctx.StatementContext)
	datas := make([][]*point, b.N)
	for i := 0; i < b.N; i++ {
		cpdata := make([]*point, 10000)
		copy(cpdata, pdata3)
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

func BenchmarkArraySort_Int(b *testing.B) {
	sc := new(stmtctx.StatementContext)
	datas := make([][]point, b.N)
	for i := 0; i < b.N; i++ {
		cdata := make([]point, 10000)
		copy(cdata, data3)
		datas[i] = cdata
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sorter := &pointSorter2{points: datas[i], sc: sc}
		sort.Sort(sorter)
	}
}

func BenchmarkHeapSort_Int(b *testing.B) {
	sc := new(stmtctx.StatementContext)
	datas := make([][]*point, b.N)
	for i := 0; i < b.N; i++ {
		cpdata := make([]*point, 10000)
		copy(cpdata, pdata3)
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

func BenchmarkPointerArraySort_Decimal(b *testing.B) {
	sc := new(stmtctx.StatementContext)
	datas := make([][]*point, b.N)
	for i := 0; i < b.N; i++ {
		cpdata := make([]*point, 10000)
		copy(cpdata, pdata4)
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

func BenchmarkArraySort_Decimal(b *testing.B) {
	sc := new(stmtctx.StatementContext)
	datas := make([][]point, b.N)
	for i := 0; i < b.N; i++ {
		cdata := make([]point, 10000)
		copy(cdata, data4)
		datas[i] = cdata
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sorter := &pointSorter2{points: datas[i], sc: sc}
		sort.Sort(sorter)
	}
}

func BenchmarkHeapSort_Decimal(b *testing.B) {
	sc := new(stmtctx.StatementContext)
	datas := make([][]*point, b.N)
	for i := 0; i < b.N; i++ {
		cpdata := make([]*point, 10000)
		copy(cpdata, pdata4)
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

func BenchmarkPointerArraySort_StringUtf8Mb4_BadData(b *testing.B) {
	sc := new(stmtctx.StatementContext)
	datas := make([][]*point, b.N)
	for i := 0; i < b.N; i++ {
		cpdata := make([]*point, 10000)
		copy(cpdata, pdata5)
		datas[i] = cpdata
	}

	var sorter *pointSorter1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sorter = &pointSorter1{points: datas[i], sc: sc}
		sorter.lessCount = 0
		sort.Sort(sorter)
	}
}

func BenchmarkHeapSort_StringUtf8Mb4_BadData(b *testing.B) {
	sc := new(stmtctx.StatementContext)
	datas := make([][]*point, b.N)
	for i := 0; i < b.N; i++ {
		cpdata := make([]*point, 10000)
		copy(cpdata, pdata5)
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
