package medium

import (
	"fmt"
	"testing"

	"leetcode/binary_tree/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/path-sum-ii/description/

// method 1 recursive DFS (top-down) Preorder Traversal
// 1) if root is nil, return result
// 2) call pathSum1DFS(root, targetSum, &path, &result)
// 3) path is a slice to store the path from root to leaf
// 4) result is a slice to store all paths
// 5) return result
// TC = O(N), SC = O(N), N is the height of tree
func pathSum1[T int](root *util.TreeNode[T], targetSum T) [][]T {
	result := [][]T{}
	if root == nil {
		return result
	}

	path := []T{}
	pathSum1DFS(root, targetSum, &path, &result)

	return result
}

func pathSum1DFS[T int](node *util.TreeNode[T], targetSum T, path *[]T, result *[][]T) {
	if node == nil {
		return
	}

	*path = append(*path, node.Val)

	if node.Left == nil && node.Right == nil && node.Val == targetSum {
		// curPath := make([]T, len(*path))
        // copy(curPath, *path)
        // *result = append(*result, curPath)

		// * key point: append a copy of ptr path into result
		*result = append(*result, append([]T{}, *path...))
	}

	pathSum1DFS(node.Left, targetSum-node.Val, path, result)
	pathSum1DFS(node.Right, targetSum-node.Val, path, result)

	// backtrack, remove the last element
	*path = (*path)[:len(*path)-1]
}

func Test_pathSum1(t *testing.T) {
	type args struct {
		root      *util.TreeNode[int]
		targetSum int
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
				root:      util.ArrayToBinaryTree([]int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, 5, 1}),
				targetSum: 22,
			},
			expected: expected{
				result: [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
			},
		},
		{
			name: "2",
			args: args{
				root:      util.ArrayToBinaryTree([]int{1, 2, 3}),
				targetSum: 5,
			},
			expected: expected{
				result: [][]int{},
			},
		},
		{
			name: "3",
			args: args{
				root:      util.ArrayToBinaryTree([]int{1, 2}),
				targetSum: 0,
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
			pathSum1(tc.args.root, tc.args.targetSum),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
