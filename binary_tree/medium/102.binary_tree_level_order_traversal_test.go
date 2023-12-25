package medium

import (
	"fmt"
	"testing"

	"leetcode/binary_tree/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/binary-tree-level-order-traversal/description/

// method 1 iterative BFS (top-down)
// 1) push root into queue
// 2) iterate queue while queue is not empty
// 3) iterate levelSize times, pop node from queue, push left and right of node into queue
// 4) append levelVals into result
// 5) return result
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
func levelOrder1[T int](root *util.TreeNode[T]) [][]T {
	result := [][]T{}
	if root == nil {
		return result
	}

	queue := []*util.TreeNode[T]{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		levelVals := []T{}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			levelVals = append(levelVals, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, levelVals)
	}

	return result
}

// method 2 recursive DFS (top-down) Preorder Traversal, if the height of tree is large, it will cause stack overflow, so it is not recommended
// 1) if root is nil, return result
// 2) call levelOrder2DFS(root, 0, &result)
// 3) return result
// TC = O(N), SC = O(N)
func levelOrder2[T int](root *util.TreeNode[T]) [][]T {
	result := [][]T{}
	if root == nil {
		return result
	}

	levelOrder2DFS(root, 0, &result)

	return result
}

// * the level of root represents the index of result
func levelOrder2DFS[T int](root *util.TreeNode[T], level int, result *[][]T) {
	if root == nil {
		return
	}

	// * this is the key point, initialize memory of index in result to append levelVals
	if len(*result) == level {
		*result = append(*result, []T{})
	}

	// * must use parentheses to wrap (*result), then operator index like (*result)[level]
	(*result)[level] = append((*result)[level], root.Val)

	levelOrder2DFS(root.Left, level+1, result)
	levelOrder2DFS(root.Right, level+1, result)
}

func Test_levelOrder1(t *testing.T) {
	type args struct {
		root *util.TreeNode[int]
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
				root: util.ArrayToBinaryTree([]int{3, 9, 20, -1, -1, 15, 7}),
			},
			expected: expected{
				result: [][]int{{3}, {9, 20}, {15, 7}},
			},
		},
		{
			name: "2",
			args: args{
				root: util.ArrayToBinaryTree([]int{1}),
			},
			expected: expected{
				result: [][]int{{1}},
			},
		},
		{
			name: "3",
			args: args{
				root: util.ArrayToBinaryTree([]int{}),
			},
			expected: expected{
				result: [][]int{},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			levelOrder1(tc.args.root),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_levelOrder2(t *testing.T) {
	type args struct {
		root *util.TreeNode[int]
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
				root: util.ArrayToBinaryTree([]int{3, 9, 20, -1, -1, 15, 7}),
			},
			expected: expected{
				result: [][]int{{3}, {9, 20}, {15, 7}},
			},
		},
		{
			name: "2",
			args: args{
				root: util.ArrayToBinaryTree([]int{1}),
			},
			expected: expected{
				result: [][]int{{1}},
			},
		},
		{
			name: "3",
			args: args{
				root: util.ArrayToBinaryTree([]int{}),
			},
			expected: expected{
				result: [][]int{},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			levelOrder2(tc.args.root),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_levelOrder1(b *testing.B) {
	root := util.ArrayToBinaryTree([]int{3, 9, 20, -1, -1, 15, 7})
	for i := 0; i < b.N; i++ {
		levelOrder1(root)
	}
}

func Benchmark_levelOrder2(b *testing.B) {
	root := util.ArrayToBinaryTree([]int{3, 9, 20, -1, -1, 15, 7})
	for i := 0; i < b.N; i++ {
		levelOrder2(root)
	}
}