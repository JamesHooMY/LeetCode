package medium

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/k-closest-points-to-origin/description/

type Point struct {
	x        int
	y        int
	distance int
}

type MinHeapPoint []Point

func (h *MinHeapPoint) Len() int { return len(*h) }

// min heap
func (h *MinHeapPoint) Less(i, j int) bool { return (*h)[i].distance < (*h)[j].distance }

func (h *MinHeapPoint) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MinHeapPoint) Push(x any) { *h = append(*h, x.(Point)) }

func (h *MinHeapPoint) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// method 1 min heap
// 1) create a min heap slice
// 2) push all points to the heap slice
// 3) init the min heap slice to a min heap using heap.Init()
// 4) pop k points from the min heap
// TC: O(NlogN), SC: O(N), N is the length of points
func kClosest1(points [][]int, k int) [][]int {
	// every time heap push will compare with the parent node, it slower than code below
	// h := &MinHeapPoint{}
	// for _, point := range points {
	// 	heap.Push(h, Point{
	// 		x:        point[0],
	// 		y:        point[1],
	// 		distance: point[0]*point[0] + point[1]*point[1],
	// 	})
	// }

	h := &MinHeapPoint{}
	for _, point := range points {
		point := Point{
			x:        point[0],
			y:        point[1],
			distance: point[0]*point[0] + point[1]*point[1],
		}
		*h = append(*h, point)
	}

	heap.Init(h) // create a min heap

	var result [][]int
	for i := 0; i < k; i++ {
		point := heap.Pop(h).(Point)
		result = append(result, []int{point.x, point.y})
	}

	return result
}

type MaxHeapPoint []Point

func (h *MaxHeapPoint) Len() int { return len(*h) }

func (h *MaxHeapPoint) Less(i, j int) bool { return (*h)[i].distance > (*h)[j].distance }

func (h *MaxHeapPoint) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MaxHeapPoint) Push(x any) { *h = append(*h, x.(Point)) }

func (h *MaxHeapPoint) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// method 2 max heap
// 1) create a max heap slice
// 2) push all points to the heap slice
// 3) pop k points from the max heap
// TC: O(NlogK), SC: O(K), N is the length of points
// * this is the best solution for me currently
func kClosest2(points [][]int, k int) [][]int {
	h := &MaxHeapPoint{}
	for _, point := range points {
		heap.Push(h, Point{
			x:        point[0],
			y:        point[1],
			distance: point[0]*point[0] + point[1]*point[1],
		})

		// keep the size of heap is k
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	var result [][]int
	for h.Len() > 0 {
		point := heap.Pop(h).(Point)
		result = append(result, []int{point.x, point.y})
	}

	return result
}

func Test_kClosest1(t *testing.T) {
	type args struct {
		points [][]int
		k      int
	}
	type expected struct {
		result [][]int
	}
	type testCase struct {
		name     string
		args     args
		expected expected
	}

	testCases := []testCase{
		{
			name: "1",
			args: args{
				points: [][]int{
					{1, 3},
					{-2, 2},
				},
				k: 1,
			},
			expected: expected{
				result: [][]int{
					{-2, 2},
				},
			},
		},
		{
			name: "2",
			args: args{
				points: [][]int{
					{3, 3},
					{5, -1},
					{-2, 4},
				},
				k: 2,
			},
			expected: expected{
				result: [][]int{
					{3, 3},
					{-2, 4},
				},
			},
		},
	}

	for _, tc := range testCases {
		assert.ElementsMatch(
			t,
			tc.expected.result,
			kClosest1(tc.args.points, tc.args.k),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_kClosest2(t *testing.T) {
	type args struct {
		points [][]int
		k      int
	}
	type expected struct {
		result [][]int
	}
	type testCase struct {
		name     string
		args     args
		expected expected
	}

	testCases := []testCase{
		{
			name: "1",
			args: args{
				points: [][]int{
					{1, 3},
					{-2, 2},
				},
				k: 1,
			},
			expected: expected{
				result: [][]int{
					{-2, 2},
				},
			},
		},
		{
			name: "2",
			args: args{
				points: [][]int{
					{3, 3},
					{5, -1},
					{-2, 4},
				},
				k: 2,
			},
			expected: expected{
				result: [][]int{
					{3, 3},
					{-2, 4},
				},
			},
		},
	}

	for _, tc := range testCases {
		assert.ElementsMatch(
			t,
			tc.expected.result,
			kClosest2(tc.args.points, tc.args.k),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_kClosest1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kClosest1([][]int{
			{1, 3},
			{-2, 2},
		}, 1)
	}
}

func Benchmark_kClosest2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kClosest2([][]int{
			{1, 3},
			{-2, 2},
		}, 1)
	}
}
