package easy

import (
	"fmt"
	"testing"

	"leetcode/binary_tree/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-depth-of-binary-tree/description/

// method 1 recursive DFS (top-down)
// getHeight same as binary_tree/easy/110.balanced_binary_tree.go
// TC = O(N), SC = O(N), N is the height of tree
// * this is the best solution for me currently
func maxDepth1[T int](root *util.TreeNode[T]) int {
	return getHeight(root)
}

// method 2 iterative BFS (top-down)
// 1) push root into queue
// 2) iterate queue while queue is not empty, pop node from queue, push left and right of node into queue
// 3) increase depth by 1
// 4) return depth
// TC = O(N), SC = O(N), N is the height of tree
func maxDepth2[T int](root *util.TreeNode[T]) int {
	if root == nil {
		return 0
	}

	queue := []*util.TreeNode[T]{root}
	depth := 0

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		depth++
	}

	return depth
}

func Test_maxDepth1(t *testing.T) {
	type args struct {
		root *util.TreeNode[int]
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
				root: util.ArrayToBinaryTree([]int{3, 9, 20, -1, -1, 15, 7}),
			},
			expected: expected{
				result: 3,
			},
		},
		{
			name: "2",
			args: args{
				root: util.ArrayToBinaryTree([]int{1, 2, 2, 3, 3, -1, -1, 4, 4}),
			},
			expected: expected{
				result: 4,
			},
		},
		{
			name: "3",
			args: args{
				root: util.ArrayToBinaryTree([]int{1}),
			},
			expected: expected{
				result: 1,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			maxDepth1(tc.args.root),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_maxDepth2(t *testing.T) {
	type args struct {
		root *util.TreeNode[int]
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
				root: util.ArrayToBinaryTree([]int{3, 9, 20, -1, -1, 15, 7}),
			},
			expected: expected{
				result: 3,
			},
		},
		{
			name: "2",
			args: args{
				root: util.ArrayToBinaryTree([]int{1, 2, 2, 3, 3, -1, -1, 4, 4}),
			},
			expected: expected{
				result: 4,
			},
		},
		{
			name: "3",
			args: args{
				root: util.ArrayToBinaryTree([]int{1}),
			},
			expected: expected{
				result: 1,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			maxDepth2(tc.args.root),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_maxDepth1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxDepth1(util.ArrayToBinaryTree([]int{3, 9, 20, -1, -1, 15, 7}))
	}
}

func Benchmark_maxDepth2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxDepth2(util.ArrayToBinaryTree([]int{3, 9, 20, -1, -1, 15, 7}))
	}
}