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

type MinHeap []Point

func (h *MinHeap) Len() int { return len(*h) }

// min heap
func (h *MinHeap) Less(i, j int) bool { return (*h)[i].distance < (*h)[j].distance }

// max heap
// func (h *MinHeap) Less(i, j int) bool { return (*h)[i].distance > (*h)[j].distance }

func (h *MinHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MinHeap) Push(x any) { *h = append(*h, x.(Point)) }

func (h *MinHeap) Pop() any {
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
// * this is the best solution for me currently
func kClosest1(points [][]int, k int) [][]int {
	// every time heap push will compare with the parent node, it slower than code below
	// h := &MinHeap{}
	// for _, point := range points {
	// 	heap.Push(h, Point{
	// 		x:        point[0],
	// 		y:        point[1],
	// 		distance: point[0]*point[0] + point[1]*point[1],
	// 	})
	// }

	h := &MinHeap{}
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
		assert.Equal(
			t,
			tc.expected.result,
			kClosest1(tc.args.points, tc.args.k),
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
