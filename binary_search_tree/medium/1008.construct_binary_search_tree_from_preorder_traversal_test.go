package medium

import (
	"testing"

	"leetcode/binary_tree/util"

	"github.com/stretchr/testify/assert"
)

// method 1 recursive DFS (top-down) Inorder Traversal
// 1) the first element of preorder is the root node
// 2) the left subtree of the root node is the nodes that are less than the root node
// 3) the right subtree of the root node is the nodes that are greater than the root node
// 4) we can use recursive to insert the node to the binary search tree
// TC: O(N), SC: O(N) for recursive stack
func bstFromPreorder1[T int](preorder []T) *util.TreeNode[T] {
	root := &util.TreeNode[T]{Val: preorder[0]}

	for _, value := range preorder[1:] {
		insertNode(root, value)
	}

	return root
}

func insertNode[T int](root *util.TreeNode[T], value T) *util.TreeNode[T] {
	if root == nil {
		return &util.TreeNode[T]{Val: value}
	}

	if value < root.Val {
		root.Left = insertNode(root.Left, value)
	} else {
		root.Right = insertNode(root.Right, value)
	}

	return root
}

func Test_bstFromPreorder1(t *testing.T) {
	type args struct {
		preorder []int
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
				preorder: []int{8, 5, 1, 7, 10, 12},
			},
			expected: expected{
				result: util.ArrayToBinaryTree([]int{8, 5, 10, 1, 7, -1, 12}),
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected.result, bstFromPreorder1(tc.args.preorder))
	}
}
