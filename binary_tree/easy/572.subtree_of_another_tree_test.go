package easy

import (
	"fmt"
	"testing"

	"leetcode/binary_tree/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/subtree-of-another-tree/

// method 1 recursive DFS (top-down) Preorder Traversal
// 1) check if root and subRoot are nil, if yes, return true
// 2) check if root and subRoot are not nil, if yes, check if root.Val == subRoot.Val, if yes, check if root and subRoot are same tree
// 3) return isSubtree1(root.Left, subRoot) || isSubtree1(root.Right, subRoot)
// TC = O(N), SC = O(N), N is the height of tree
// * this is the best solution for me currently
func isSubtree1[T int](root *util.TreeNode[T], subRoot *util.TreeNode[T]) bool {
	if root == nil || subRoot == nil {
		return false
	}

	if root.Val == subRoot.Val && isSameTree1(root, subRoot) {
		return true
	}

	return isSubtree1(root.Left, subRoot) || isSubtree1(root.Right, subRoot)
}

func Test_isSubtree1(t *testing.T) {
	type args struct {
		root    *util.TreeNode[int]
		subRoot *util.TreeNode[int]
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
				root:    util.ArrayToBinaryTree([]int{3, 4, 5, 1, 2}),
				subRoot: util.ArrayToBinaryTree([]int{4, 1, 2}),
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				root:    util.ArrayToBinaryTree([]int{3, 4, 5, 1, 2, -1, -1, -1, -1, 0}),
				subRoot: util.ArrayToBinaryTree([]int{4, 1, 2}),
			},
			expected: expected{
				result: false,
			},
		},
		{
			name: "3",
			args: args{
				root:    util.ArrayToBinaryTree([]int{1, 1}),
				subRoot: util.ArrayToBinaryTree([]int{1}),
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "4",
			args: args{
				root:    util.ArrayToBinaryTree([]int{1,-1,1,-1,1,-1,1,-1,1,-1,1,-1,1,-1,1,-1,1,-1,1,-1,1,2}),
				subRoot: util.ArrayToBinaryTree([]int{1,-1,1,-1,1,-1,1,-1,1,-1,1,2}),
			},
			expected: expected{
				result: true,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			isSubtree1(tc.args.root, tc.args.subRoot),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
