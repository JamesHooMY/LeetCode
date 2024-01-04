package medium

import (
	"fmt"
	"testing"

	"leetcode/binary_tree/util"

	"github.com/stretchr/testify/assert"
)

// https://www.cnblogs.com/grandyang/p/5306162.html

// method 1 recursive DFS (top-down) Inorder Traversal
// 1) inorder traversal rule is left -> root -> right
// 2) if root is nil, return nil
// 3) if root.Val > target.Val mean the target is in the left subtree, current root is a candidate as successor, so we record it, then keep going left
// 4) if root.Val <= target.Val mean the target is in the right subtree, so we just keep going right
// 5) return successor
// TC: O(N), SC: O(1)
// * this is the best solution for me currently
func inorderSuccessor1[T int](root *util.TreeNode[T], target *util.TreeNode[T]) *util.TreeNode[T] {
	var successor *util.TreeNode[T]

	for root != nil {
		if root.Val > target.Val {
			// * this is the key point, we need to record the successor, then keep going left
			successor = root
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return successor
}

func Test_inorderSuccessor1(t *testing.T) {
	type args struct {
		root   *util.TreeNode[int]
		target *util.TreeNode[int]
	}
	type expected struct {
		result *util.TreeNode[int]
	}
	type testCase struct {
		name     string
		args     args
		expected expected
	}

	roots := []*util.TreeNode[int]{
		util.ArrayToBinaryTree([]int{2, 1, 3}),
		util.ArrayToBinaryTree([]int{5, 3, 6, 2, 4, -1, -1, 1}),
	}

	testCases := []testCase{
		{
			name: "1",
			args: args{
				root:   roots[0],
				target: roots[0].Left,
			},
			expected: expected{
				result: roots[0],
			},
		},
		{
			name: "2",
			args: args{
				root:   roots[1],
				target: roots[1].Right,
			},
			expected: expected{
				result: nil,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			inorderSuccessor1(tc.args.root, tc.args.target),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
