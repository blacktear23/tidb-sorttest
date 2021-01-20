package main

import (
	"container/heap"
	"encoding/hex"
	"fmt"
	"math/rand"
	"sort"
	"time"

	"github.com/pingcap/tidb/sessionctx/stmtctx"
	"github.com/pingcap/tidb/types"
)

type point struct {
	value types.Datum
}

func rangePointLess1(a, b *point, sc *stmtctx.StatementContext) bool {
	cmp, err := a.value.CompareDatum(sc, &b.value)
	if err != nil {
		return true
	}
	return cmp < 0
}

func rangePointLess2(a, b point, sc *stmtctx.StatementContext) bool {
	cmp, err := a.value.CompareDatum(sc, &b.value)
	if err != nil {
		return true
	}
	return cmp < 0
}

type pointSorter1 struct {
	points    []*point
	sc        *stmtctx.StatementContext
	lessCount int
}

func (r *pointSorter1) Len() int {
	return len(r.points)
}

func (r *pointSorter1) Less(i, j int) bool {
	a := r.points[i]
	b := r.points[j]
	r.lessCount++
	return rangePointLess1(a, b, r.sc)
}

func (r *pointSorter1) Swap(i, j int) {
	r.points[i], r.points[j] = r.points[j], r.points[i]
}

type pointSorter2 struct {
	points []point
	sc     *stmtctx.StatementContext
}

func (r *pointSorter2) Len() int {
	return len(r.points)
}

func (r *pointSorter2) Less(i, j int) bool {
	a := r.points[i]
	b := r.points[j]
	return rangePointLess2(a, b, r.sc)
}

func (r *pointSorter2) Swap(i, j int) {
	r.points[i], r.points[j] = r.points[j], r.points[i]
}

func getRandomString() string {
	data := make([]byte, 10)
	rand.Read(data)
	return hex.EncodeToString(data)
}

func getRandomInt() uint32 {
	return rand.Uint32()
}

func prepareIntArrays(n int) ([]*point, []point) {
	ret1 := make([]*point, n)
	ret2 := make([]point, n)
	for i := 0; i < n; i++ {
		d := types.NewDatum(getRandomInt())
		ret1[i] = &point{value: d}
		ret2[i] = point{value: d}
	}
	return ret1, ret2
}

func prepareStringArrays(n int, collation string) ([]*point, []point) {
	ret1 := make([]*point, n)
	ret2 := make([]point, n)
	for i := 0; i < n; i++ {
		d := types.NewDatum(getRandomString())
		d.SetCollation(collation)
		ret1[i] = &point{value: d}
		ret2[i] = point{value: d}
	}
	return ret1, ret2
}

type PointHeap struct {
	points    []*point
	sc        *stmtctx.StatementContext
	lessCount int
}

func (h *PointHeap) Len() int {
	return len(h.points)
}

func (h *PointHeap) Less(i, j int) bool {
	a := h.points[i]
	b := h.points[j]
	h.lessCount++
	return rangePointLess1(a, b, h.sc)
}

func (h *PointHeap) Swap(i, j int) {
	h.points[i], h.points[j] = h.points[j], h.points[i]
}

func (h *PointHeap) Push(x interface{}) {
	h.points = append(h.points, x.(*point))
}

func (h *PointHeap) Pop() interface{} {
	old := h.points
	n := len(old)
	x := old[n-1]
	h.points = old[0 : n-1]
	return x
}

func heapSort(h *PointHeap, data []*point) (*PointHeap, []*point) {
	heap.Init(h)
	for _, item := range data {
		heap.Push(h, item)
	}
	return h, h.points
}

func main() {
	pdata, data := prepareStringArrays(10000, "utf8mb4_general_ci")
	sc := new(stmtctx.StatementContext)

	start := time.Now()
	sorter1 := &pointSorter1{points: pdata, sc: sc}
	sort.Sort(sorter1)
	fmt.Println("Pointer Array Sort Use Time", time.Now().Sub(start))

	start = time.Now()
	sorter2 := &pointSorter2{points: data, sc: sc}
	sort.Sort(sorter2)
	fmt.Println("Array Sort Use Time", time.Now().Sub(start))
}
