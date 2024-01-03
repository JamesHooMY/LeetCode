package medium

import (
	"fmt"
	"testing"

	"leetcode/binary_tree/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/

// method 1 recursive DFS (top-down) Inorder Traversal
// 1) if root is nil, return -1
// 2) call kthSmallest1DFS(root.Left, k)
// 3) count++
// 4) if count == k, return root.Val
// 5) call kthSmallest1DFS(root.Right, k)
// 6) return -1
// TC: O(N), SC: O(N)
// * this is the best solution for me currently
func kthSmallest1[T int](root *util.TreeNode[T], k T) T {
	count := 0
	return kthSmallest1DFS(root, k, &count)
}

func kthSmallest1DFS[T int](root *util.TreeNode[T], k T, count *int) T {
	if root == nil {
		return -1
	}

	leftVal := kthSmallest1DFS(root.Left, k, count)
	if leftVal != -1 {
		return leftVal
	}

	*count++
	if *count == int(k) {
		return root.Val
	}

	return kthSmallest1DFS(root.Right, k, count)
}

// method 2 stack iterative DFS (top-down) Inorder Traversal
// 1) create stack
// 2) create cur = root
// 3) for cur != nil || len(stack) > 0
// 4) if cur != nil, push cur to stack, cur = cur.Left
// 5) else
// 6) pop top from stack, count++
// 7) if count == k, return top.Val
// 8) cur = top.Right
// 9) return -1
// TC: O(N), SC: O(N)
func kthSmallest2[T int](root *util.TreeNode[T], k T) T {
	stack := []*util.TreeNode[T]{}

	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]   // stack top
		stack = stack[:len(stack)-1] // pop stack top

		k--
		if k == 0 {
			return root.Val
		}

		root = root.Right
	}
}

func Test_kthSmallest1(t *testing.T) {
	type args struct {
		root *util.TreeNode[int]
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
				root: util.ArrayToBinaryTree([]int{3, 1, 4, -1, 2}),
				k:    1,
			},
			expected: expected{
				result: 1,
			},
		},
		{
			name: "2",
			args: args{
				root: util.ArrayToBinaryTree([]int{5, 3, 6, 2, 4, -1, -1, 1}),
				k:    3,
			},
			expected: expected{
				result: 3,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			kthSmallest1(tc.args.root, tc.args.k),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_kthSmallest2(t *testing.T) {
	type args struct {
		root *util.TreeNode[int]
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
				root: util.ArrayToBinaryTree([]int{3, 1, 4, -1, 2}),
				k:    1,
			},
			expected: expected{
				result: 1,
			},
		},
		{
			name: "2",
			args: args{
				root: util.ArrayToBinaryTree([]int{5, 3, 6, 2, 4, -1, -1, 1}),
				k:    3,
			},
			expected: expected{
				result: 3,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			kthSmallest2(tc.args.root, tc.args.k),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_kthSmallest1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kthSmallest1(util.ArrayToBinaryTree([]int{3, 1, 4, -1, 2}), 1)
	}
}

func Benchmark_kthSmallest2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kthSmallest2(util.ArrayToBinaryTree([]int{3, 1, 4, -1, 2}), 1)
	}
}