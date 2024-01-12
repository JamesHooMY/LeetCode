package medium

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"

	commonUtil "leetcode/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-k-closest-elements/description/

type Pair struct {
	value, diff int
}

type MaxHeapPair []Pair

func (h *MaxHeapPair) Len() int { return len(*h) }

func (h *MaxHeapPair) Less(i, j int) bool {
	if (*h)[i].diff == (*h)[j].diff {
		// |a - x| == |b - x| and a < b
		return (*h)[i].value > (*h)[j].value
	}

	return (*h)[i].diff > (*h)[j].diff
}

func (h *MaxHeapPair) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MaxHeapPair) Push(x any) { *h = append(*h, x.(Pair)) }

func (h *MaxHeapPair) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// method 1 max heap this method is contrast to min heap in 692.top_k_frequent_words_test.go
// 1) use max heap to store the pair of (value, |value - x|)
// 2) iterate the array, push the pair into the heap, and keep the size of heap is k
// 3) pop the top k pairs from the heap to get the result
// 4) sort the result
// TC: O(NlogK), SC: O(K), N is the length of array, K is the size of heap
func findClosestElements1(arr []int, k int, x int) []int {
	h := &MaxHeapPair{}
	// heap.Init(h)

	for _, v := range arr {
		heap.Push(h, Pair{v, commonUtil.Abs(v - x)})

		// keep the size of heap is k
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := []int{}
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(Pair).value)
	}

	sort.Ints(res)
	return res
}

// method 2 sort
// 1) sort the array by the absolute difference between the element and x
// 2) sort the first k elements
// 3) sort the result
// TC: O(NlogN), SC: O(1)
// * this is the best solution for me currently
func findClosestElements2(arr []int, k int, x int) []int {
	sort.Slice(arr, func(i, j int) bool {
		if commonUtil.Abs(arr[i]-x) == commonUtil.Abs(arr[j]-x) {
			return arr[i] < arr[j]
		}

		return commonUtil.Abs(arr[i]-x) < commonUtil.Abs(arr[j]-x)
	})

	sort.Ints(arr[:k])

	return arr[:k]
}

func Test_findClosestElements1(t *testing.T) {
	type args struct {
		arr []int
		k   int
		x   int
	}
	type expected struct {
		result []int
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
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   3,
			},
			expected: expected{
				result: []int{1, 2, 3, 4},
			},
		},
		{
			name: "2",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   -1,
			},
			expected: expected{
				result: []int{1, 2, 3, 4},
			},
		},
		{
			name: "3",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   6,
			},
			expected: expected{
				result: []int{2, 3, 4, 5},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			findClosestElements1(tc.args.arr, tc.args.k, tc.args.x),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_findClosestElements2(t *testing.T) {
	type args struct {
		arr []int
		k   int
		x   int
	}
	type expected struct {
		result []int
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
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   3,
			},
			expected: expected{
				result: []int{1, 2, 3, 4},
			},
		},
		{
			name: "2",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   -1,
			},
			expected: expected{
				result: []int{1, 2, 3, 4},
			},
		},
		{
			name: "3",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   6,
			},
			expected: expected{
				result: []int{2, 3, 4, 5},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			findClosestElements2(tc.args.arr, tc.args.k, tc.args.x),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_findClosestElements1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findClosestElements1([]int{1, 2, 3, 4, 5}, 4, 3)
	}
}

func Benchmark_findClosestElements2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findClosestElements2([]int{1, 2, 3, 4, 5}, 4, 3)
	}
}