package easy

import (
	"fmt"
	"testing"

	"leetcode/binary_tree/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/invert-binary-tree/

// method 1 recursive DFS (top-down)
// 1) invert left subtree, invert right subtree
// 2) swap left and right
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
func invertTree1[T int](root *util.TreeNode[T]) *util.TreeNode[T] {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	// root.Left, root.Right = root.Right, root.Left
	// invertTree1(root.Left)
	// invertTree1(root.Right)
	root.Left, root.Right = invertTree1(root.Right), invertTree1(root.Left)

	return root
}

func Test_invertTree1(t *testing.T) {
	type args struct {
		root *util.TreeNode[int]
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
				root: util.ArrayToBinaryTree([]int{4, 2, 7, 1, 3, 6, 9}),
			},
			expected: expected{
				result: util.ArrayToBinaryTree([]int{4, 7, 2, 9, 6, 3, 1}),
			},
		},
		{
			name: "2",
			args: args{
				root: util.ArrayToBinaryTree([]int{1, 2}),
			},
			expected: expected{
				result: util.ArrayToBinaryTree([]int{1, -1, 2}),
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			invertTree1(tc.args.root),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
