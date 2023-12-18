package easy

import (
	"fmt"
	"testing"

	"leetcode/tree/util"
	commonUtil "leetcode/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/balanced-binary-tree/description/

// method 1 recursive DFS (top-down)
// 1) get height of left subtree, get height of right subtree
// 2) if |height of left subtree - height of right subtree| <= 1, then check if left subtree is balanced, check if right subtree is balanced
// 3) return true if both left subtree and right subtree are balanced
// TC = O(N^2), SC = O(N), N is the height of tree
func isBalanced1[T any](root *util.TreeNode[T]) bool {
	// child nodes (Left, Right) of leaf nodes are nil, thus return true directly
	if root == nil {
		return true
	}

	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)

	// commonUtil.Abs(leftHeight-rightHeight) of each nodes <= 1 means balanced
	if commonUtil.Abs(leftHeight-rightHeight) > 1 {
		return false
	}

	// isBalanced1 here will cause the TC from O(N) to O(N^2), due to the getHeight will be called for each node repeatedly !!!
	return isBalanced1(root.Left) && isBalanced1(root.Right)
}

func getHeight[T any](node *util.TreeNode[T]) int {
	// child nodes (Left, Right) of leaf nodes are nil, thus return 0 directly
	if node == nil {
		return 0
	}

	// leaf nodes will get leftHeight and rightHeight are equal to 0
	leftHeight := getHeight(node.Left)
	rightHeight := getHeight(node.Right)


	// return height of each parent node
	return commonUtil.Max(leftHeight, rightHeight) + 1
}

// method 2 recursive DFS (top-down) dynamic programming
// 1) get height of left subtree, get height of right subtree
// 2) if height of left subtree or height of right subtree is -1, return -1
// 3) if |height of left subtree - height of right subtree| > 1, return -1
// 4) return max(height of left subtree, height of right subtree) + 1
// TC = O(N), SC = O(N), N is the height of tree
// * this is the best solution for me currently
func isBalanced2[T any](root *util.TreeNode[T]) bool {
	return getHeightAndBalance(root) != -1
}

// return -1 means the tree is not balanced, otherwise return the height of parent node
func getHeightAndBalance[T any](node *util.TreeNode[T]) int {
	// child nodes (Left, Right) of leaf nodes are nil, thus return 0 directly
	if node == nil {
		return 0
	}

	leftHeight := getHeightAndBalance(node.Left)
	// if node not balance then return -1 directly, skip below process
	if leftHeight == -1 {
		return -1
	}

	rightHeight := getHeightAndBalance(node.Right)
	// if node not balance then return -1 directly, skip below process
	if rightHeight == -1 {
		return -1
	}

	// commonUtil.Abs(leftHeight-rightHeight) of each nodes <= 1 means balanced, check each node whether it is balance !!!
	if commonUtil.Abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	// return height of each parent node
	return commonUtil.Max(leftHeight, rightHeight) + 1
}

func Test_isBalanced1(t *testing.T) {
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
				root: util.ArrayToBinaryTree([]int{3, 9, 20, -1, -1, 15, 7}),
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				root: util.ArrayToBinaryTree([]int{1, 2, 2, 3, 3, -1, -1, 4, 4}),
			},
			expected: expected{
				result: false,
			},
		},
		{
			name: "3",
			args: args{
				root: util.ArrayToBinaryTree([]int{}),
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
			isBalanced1(tc.args.root),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_isBalanced2(t *testing.T) {
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
				root: util.ArrayToBinaryTree([]int{3, 9, 20, -1, -1, 15, 7}),
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				root: util.ArrayToBinaryTree([]int{1, 2, 2, 3, 3, -1, -1, 4, 4}),
			},
			expected: expected{
				result: false,
			},
		},
		{
			name: "3",
			args: args{
				root: util.ArrayToBinaryTree([]int{}),
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
			isBalanced2(tc.args.root),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_isBalanced1(b *testing.B) {
	root := util.ArrayToBinaryTree([]int{3, 9, 20, -1, -1, 15, 7})

	for i := 0; i < b.N; i++ {
		isBalanced1(root)
	}
}

func Benchmark_isBalanced2(b *testing.B) {
	root := util.ArrayToBinaryTree([]int{3, 9, 20, -1, -1, 15, 7})

	for i := 0; i < b.N; i++ {
		isBalanced2(root)
	}
}
