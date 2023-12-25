package medium

import (
	"fmt"
	"testing"

	"leetcode/binary_tree/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/binary-tree-right-side-view/

// method 1 iterative BFS (top-down)
// 1) push root into queue
// 2) iterate queue while queue is not empty
// 3) iterate levelSize times, pop node from queue, push left and right of node into queue
// 4) append the rightmost node of each level into result
// 5) return result
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
func rightSideView1[T int](root *util.TreeNode[T]) []T {
	result := []T{}
	if root == nil {
		return result
	}

	queue := []*util.TreeNode[T]{root}
	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// * the rightmost node of each level, rear element of queue
			if i == levelSize-1 {
				result = append(result, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return result
}

// method 2 recursive DFS (top-down)
// 1) if root is nil, return result
// 2) call rightSideView2DFS(root, 0, &result)
// 3) return result
// TC = O(N), SC = O(N)
func rightSideView2[T int](root *util.TreeNode[T]) []T {
	result := []T{}
	if root == nil {
		return result
	}

	rightSideView2DFS(root, 0, &result)

	return result
}

func rightSideView2DFS[T int](root *util.TreeNode[T], level int, result *[]T) {
	if root == nil {
		return
	}

	// * only add one rightmost node of each level into result
	if level == len(*result) {
		*result = append(*result, root.Val)
	}

	// * this is the key point, traverse right first, then left
	rightSideView2DFS(root.Right, level+1, result)
	rightSideView2DFS(root.Left, level+1, result)
}

func Test_rightSideView1(t *testing.T) {
	type args struct {
		root *util.TreeNode[int]
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
				root: util.ArrayToBinaryTree([]int{1, 2, 3, -1, 5, -1, 4}),
			},
			expected: expected{
				result: []int{1, 3, 4},
			},
		},
		{
			name: "2",
			args: args{
				root: util.ArrayToBinaryTree([]int{1, -1, 3}),
			},
			expected: expected{
				result: []int{1, 3},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			rightSideView1(tc.args.root),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_rightSideView2(t *testing.T) {
	type args struct {
		root *util.TreeNode[int]
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
				root: util.ArrayToBinaryTree([]int{1, 2, 3, -1, 5, -1, 4}),
			},
			expected: expected{
				result: []int{1, 3, 4},
			},
		},
		{
			name: "2",
			args: args{
				root: util.ArrayToBinaryTree([]int{1, -1, 3}),
			},
			expected: expected{
				result: []int{1, 3},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			rightSideView2(tc.args.root),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
