package easy

import (
	"fmt"
	"testing"

	"leetcode/binary_tree/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/

// method 1 recursive DFS (top-down) Preorder Traversal
// 1) if nums is empty, return nil
// 2) midIndex = len(nums) / 2
// 3) midVal = nums[midIndex]
// 4) create root with midVal
// 5) root.Left = sortedArrayToBST1(nums[:midIndex])
// 6) root.Right = sortedArrayToBST1(nums[midIndex+1:])
// 7) return root
// TC: O(N), SC: O(N)
func sortedArrayToBST1[T int](nums []T) *util.TreeNode[T] {
	if len(nums) == 0 {
		return nil
	}
	/*
		odd array: [-10, -3, 0, 5, 9] midIndex = 5 / 2 = 2, midVal = 0
		even array: [-10, -3, 0, 5, 9, 10] midIndex = 6 / 2 = 3, midVal = 5
	*/
	midIndex := len(nums) / 2
	root := &util.TreeNode[T]{
		Val: nums[midIndex],
	}

	root.Left = sortedArrayToBST1(nums[:midIndex])
	root.Right = sortedArrayToBST1(nums[midIndex+1:])

	return root
}

// method 2 use helper function
// 1) call sortedArrayToBSTHelper(nums, 0, len(nums)-1)
// 2) if start > end, return nil
// 3) midIndex = start + (end - start) / 2
// 4) midVal = nums[midIndex]
// 5) create root with midVal
// 6) root.Left = sortedArrayToBSTHelper(nums, start, midIndex-1)
// 7) root.Right = sortedArrayToBSTHelper(nums, midIndex+1, end)
// 8) return root
// TC: O(N), SC: O(N)
func sortedArrayToBST2[T int](nums []T) *util.TreeNode[T] {
	return sortedArrayToBSTHelper(nums, 0, len(nums)-1)
}

func sortedArrayToBSTHelper[T int](nums []T, start, end int) *util.TreeNode[T] {
	if start > end {
		return nil
	}

	/*
		odd array: [-10, -3, 0, 5, 9] midIndex = 0 + (4 - 0) / 2 = 2, midVal = 0
		even array: [-10, -3, 0, 5, 9, 10] midIndex = 0 + (5 - 0) / 2 = 2, midVal = 0
	*/
	midIndex := start + (end-start)/2
	root := &util.TreeNode[T]{
		Val: nums[midIndex],
	}

	root.Left = sortedArrayToBSTHelper(nums, start, midIndex-1)
	root.Right = sortedArrayToBSTHelper(nums, midIndex+1, end)

	return root
}

func Test_sortedArrayToBST1(t *testing.T) {
	type args struct {
		nums []int
	}
	type expected struct {
		result *util.TreeNode[int]
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
				nums: []int{-10, -3, 0, 5, 9},
			},
			expected: expected{
				result: util.ArrayToBinaryTree([]int{0, -3, 9, -10, -1, 5}),
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			sortedArrayToBST1(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_sortedArrayToBST2(t *testing.T) {
	type args struct {
		nums []int
	}
	type expected struct {
		result *util.TreeNode[int]
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
				nums: []int{-10, -3, 0, 5, 9},
			},
			expected: expected{
				result: util.ArrayToBinaryTree([]int{0, -3, 9, -10, -1, 5}),
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			sortedArrayToBST2(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_sortedArrayToBST1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortedArrayToBST1([]int{-10, -3, 0, 5, 9})
	}
}

func Benchmark_sortedArrayToBST2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortedArrayToBST2([]int{-10, -3, 0, 5, 9})
	}
}
