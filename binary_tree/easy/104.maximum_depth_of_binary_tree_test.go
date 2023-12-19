package easy

import (
	"fmt"
	"testing"

	"leetcode/binary_tree/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-depth-of-binary-tree/description/

// method 1 recursive DFS (top-down)
// getHeight same as binary_tree/easy/110.balanced_binary_tree.go
// TC = O(N), SC = O(N), N is the height of tree
func maxDepth1[T any](root *util.TreeNode[T]) int {
	return getHeight(root)
}

func Test_maxDepth1(t *testing.T) {
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
				root: util.ArrayToBinaryTree([]int{3, 9, 20, -1, -1, 15, 7}),
			},
			expected: expected{
				result: 3,
			},
		},
		{
			name: "2",
			args: args{
				root: util.ArrayToBinaryTree([]int{1, 2, 2, 3, 3, -1, -1, 4, 4}),
			},
			expected: expected{
				result: 4,
			},
		},
		{
			name: "3",
			args: args{
				root: util.ArrayToBinaryTree([]int{1}),
			},
			expected: expected{
				result: 1,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			maxDepth1(tc.args.root),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
