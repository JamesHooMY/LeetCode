package medium

import (
	"fmt"
	"testing"

	"leetcode/binary_tree/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/

// method 1
func buildTree1[T int](preorder []T, inorder []T) *util.TreeNode[T] {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	root := &util.TreeNode[T]{Val: rootVal}

	// * find the index of rootVal in inorder
	var rootIdxInorder int
	for i, v := range inorder {
		if v == rootVal {
			rootIdxInorder = i
			break
		}
	}

	// * preorder: [root, left, right]
	// * inorder: [left, root, right]
	root.Left = buildTree1(preorder[1:rootIdxInorder+1], inorder[:rootIdxInorder])
	root.Right = buildTree1(preorder[rootIdxInorder+1:], inorder[rootIdxInorder+1:])

	return root
}

func Test_buildTree1(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
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
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			expected: expected{
				result: util.ArrayToBinaryTree([]int{3, 9, 20, -1, -1, 15, 7}),
			},
		},
		{
			name: "2",
			args: args{
				preorder: []int{-1},
				inorder:  []int{-1},
			},
			expected: expected{
				result: util.ArrayToBinaryTree([]int{-1}),
			},
		},
		{
			name: "3",
			args: args{
				preorder: []int{1, 2},
				inorder:  []int{1, 2},
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
			buildTree1(tc.args.preorder, tc.args.inorder),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
