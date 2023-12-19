package easy

import (
	"fmt"
	"testing"

	"leetcode/binary_tree/util"
	commonUtil "leetcode/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/diameter-of-binary-tree/

// method 1 recursive DFS (top-down)
// 1) calculate the height of left subtree, calculate the height of right subtree
// 2) calculate the diameter of current node and update the maxDiameter
// 3) return the height of parent node
// TC = O(N), SC = O(N), N is the height of tree
// * this is the best solution for me currently
func diameterOfBinaryTree1[T any](root *util.TreeNode[T]) int {
	if root == nil {
		return 0
	}

	maxDiameter := 0
	calculateDiameter(root, &maxDiameter)

	return maxDiameter
}

func calculateDiameter[T any](root *util.TreeNode[T], maxDiameter *int) int {
	// child nodes (Left, Right) of leaf nodes are nil, thus return 0 directly
	if root == nil {
		return 0
	}

	leftHeight := calculateDiameter(root.Left, maxDiameter)
	rightHeight := calculateDiameter(root.Right, maxDiameter)

	// check if current node is the root of the longest path(diameter)
	*maxDiameter = commonUtil.Max(*maxDiameter, leftHeight+rightHeight)

	// return height of each parent node
	return commonUtil.Max(leftHeight, rightHeight) + 1
}

func Test_diameterOfBinaryTree1(t *testing.T) {
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
				root: util.ArrayToBinaryTree([]int{1, 2, 3, 4, 5}),
			},
			expected: expected{
				result: 3,
			},
		},
		{
			name: "2",
			args: args{
				root: util.ArrayToBinaryTree([]int{1, 2}),
			},
			expected: expected{
				result: 1,
			},
		},
		{
			name: "3",
			args: args{
				root: util.ArrayToBinaryTree([]int{1, 2, 3}),
			},
			expected: expected{
				result: 2,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			diameterOfBinaryTree1(tc.args.root),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
