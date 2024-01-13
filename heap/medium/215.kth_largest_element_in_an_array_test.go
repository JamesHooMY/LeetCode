package medium

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/kth-largest-element-in-an-array/description/

type MinHeapNum []int // store the number

func (h *MinHeapNum) Len() int { return len(*h) }

// max heap
func (h *MinHeapNum) Less(i, j int) bool { return (*h)[i] < (*h)[j] }

func (h *MinHeapNum) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MinHeapNum) Push(x any) { *h = append(*h, x.(int)) }

func (h *MinHeapNum) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// method 1 min heap
// 1) build a min heap to store the number
// 2) keep the size of heap to k
// 3) pop the top of heap
// TC: O(NlogK), SC: O(K), N is the length of nums, K is the size of heap
// * this is the best solution for me currently
func findKthLargest1(nums []int, k int) int {
	h := &MinHeapNum{}
	for _, num := range nums {
		heap.Push(h, num)

		// * this is the key point, keep the size of heap is k, so the top of heap is the kth largest element
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return heap.Pop(h).(int)
}

// method 2 quick select, this method TLE in leetcode, cannot pass the test
// 1) use quick select to find the kth largest element
// 2) the pivot is the kth largest element
// 3) the left side of pivot is larger than pivot, the right side of pivot is smaller or equal to pivot
// 4) if pivot is k-1, return the pivot
// 5) if pivot is larger than k-1, do quick select on the left side
// 6) if pivot is smaller than k-1, do quick select on the right side
// TC: O(N), SC: O(1), N is the length of nums
func findKthLargest2(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k)
}

func quickSelect(nums []int, leftIndex, rightIndex, k int) int {
	pivot := partition(nums, leftIndex, rightIndex)
	if pivot == k-1 {
		return nums[pivot]
	} else if pivot > k-1 {
		return quickSelect(nums, leftIndex, pivot-1, k)
	} else {
		return quickSelect(nums, pivot+1, rightIndex, k)
	}
}

func partition(nums []int, leftIndex, rightIndex int) int {
	pivot := nums[rightIndex]
	for i := leftIndex; i < rightIndex; i++ {
		// put the number which is larger than pivot to the left side
		if nums[i] > pivot {
			nums[leftIndex], nums[i] = nums[i], nums[leftIndex]
			leftIndex++
		}
	}

	// change the pivot to the position, left side is larger than pivot, rightIndex side is smaller or equal to pivot
	nums[leftIndex], nums[rightIndex] = nums[rightIndex], nums[leftIndex]

	return leftIndex
}

func Test_findKthLargest1(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	type expected struct {
		result int
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
				nums: []int{3, 2, 1, 5, 6, 4},
				k:    2,
			},
			expected: expected{
				result: 5,
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
				k:    4,
			},
			expected: expected{
				result: 4,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			findKthLargest1(tc.args.nums, tc.args.k),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_findKthLargest2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	type expected struct {
		result int
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
				nums: []int{3, 2, 1, 5, 6, 4},
				k:    2,
			},
			expected: expected{
				result: 5,
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
				k:    4,
			},
			expected: expected{
				result: 4,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			findKthLargest2(tc.args.nums, tc.args.k),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_findKthLargest1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findKthLargest1([]int{3, 2, 1, 5, 6, 4}, 2)
	}
}

func Benchmark_findKthLargest2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findKthLargest2([]int{3, 2, 1, 5, 6, 4}, 2)
	}
}