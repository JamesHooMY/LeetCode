package medium

import (
	"fmt"
	"testing"

	"leetcode/binary_tree/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/validate-binary-search-tree/description/

// method 1 recursive DFS (top-down) Inorder Traversal
func isValidBST1[T int](root *util.TreeNode[T]) bool {
	return isValidBSTHelper1(root, nil, nil)
}

func isValidBSTHelper1[T int](root, min, max *util.TreeNode[T]) bool {
	if root == nil {
		return true
	}

	// isValidBSTHelper1(root.Right, root, max), min == root, max == nil
	if min != nil && root.Val <= min.Val {
		return false
	}

	// isValidBSTHelper1(root.Left, min, root), min == nil, max == root
	if max != nil && root.Val >= max.Val {
		return false
	}

	return isValidBSTHelper1(root.Left, min, root) && isValidBSTHelper1(root.Right, root, max)
}

func Test_isValidBST1(t *testing.T) {
	type args struct {
		root *util.TreeNode[int]
	}
	type expected struct {
		result bool
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
				root: util.ArrayToBinaryTree([]int{2, 1, 3}),
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				root: util.ArrayToBinaryTree([]int{5, 1, 4, -1, -1, 3, 6}),
			},
			expected: expected{
				result: false,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			isValidBST1(tc.args.root),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
