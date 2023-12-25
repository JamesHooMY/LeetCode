package easy

import (
	"fmt"
	"testing"

	"leetcode/binary_tree/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/same-tree/description/

// method 1 recursive DFS (top-down) Postorder Traversal
// 1) normally child nodes (Left, Right) of leaf nodes are nil, thus return true directly
// 2) after the above 1) check, if one of p and q is nil, that means exist child nodes (Left, Right) not match, return false
// 3) after the above 1) and 2) check, if both p and q are not nil, check if p.Val == q.Val, if not match, return false
// 4) return isSameTree1(p.Left, q.Left) && isSameTree1(p.Right, q.Right)
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
func isSymmetric1[T int](root *util.TreeNode[T]) bool {
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)
}

func isMirror[T int](left *util.TreeNode[T], right *util.TreeNode[T]) bool {
	// 1) normally child nodes (Left, Right) of leaf nodes are nil, thus return true directly
	if left == nil && right == nil {
		return true
	}

	// 2) after the above 1) check, if one of p and q is nil, that means exist child nodes (Left, Right) not match, return false
	if left == nil || right == nil {
		return false
	}

	// 3) after the above 1) and 2) check, if both p and q are not nil, check if p.Val == q.Val, if not match, return false
	if left.Val != right.Val {
		return false
	}

	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

// method 2 iterative BFS (top-down) like binary_tree/easy/100.same_tree.go
// 1) push p and q into queue
// 2) iterate queue while queue is not empty, pop node from queue, check if p.Val == q.Val
// 3) if p.Val != q.Val, return false
// 4) if p.Left != nil, q.Left != nil, push p.Left and q.Left into queue
// 5) if p.Right != nil, q.Right != nil, push p.Right and q.Right into queue
// 6) return true
// TC = O(N), SC = O(N)
func isSymmetric2[T int](root *util.TreeNode[T]) bool {
	queue := []*util.TreeNode[T]{root.Left, root.Right}

	for len(queue) > 0 {
		leftNode, rightNode := queue[0], queue[1]
		queue = queue[2:] // pop front 2 elements

		if leftNode == nil && rightNode == nil {
			return true
		} else if leftNode == nil || rightNode == nil || leftNode.Val != rightNode.Val {
			return false
		}

		queue = append(queue, leftNode.Left, rightNode.Right, leftNode.Right, rightNode.Left)
	}

	return true
}

func Test_isSymmetric1(t *testing.T) {
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
				root: util.ArrayToBinaryTree([]int{1, 2, 2, 3, 4, 4, 3}),
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				root: util.ArrayToBinaryTree([]int{1, 2, 2, -1, 3, -1, 3}),
			},
			expected: expected{
				result: false,
			},
		},
		{
			name: "3",
			args: args{
				root: util.ArrayToBinaryTree([]int{1, 2, 2, 3, -1, -1, 3}),
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
			isSymmetric1(tc.args.root),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_isSymmetric2(t *testing.T) {
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
				root: util.ArrayToBinaryTree([]int{1, 2, 2, 3, 4, 4, 3}),
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				root: util.ArrayToBinaryTree([]int{1, 2, 2, -1, 3, -1, 3}),
			},
			expected: expected{
				result: false,
			},
		},
		{
			name: "3",
			args: args{
				root: util.ArrayToBinaryTree([]int{1, 2, 2, 3, -1, -1, 3}),
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
			isSymmetric2(tc.args.root),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_isSymmetric1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isSymmetric1(util.ArrayToBinaryTree([]int{1, 2, 2, 3, 4, 4, 3}))
	}
}

func Benchmark_isSymmetric2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isSymmetric2(util.ArrayToBinaryTree([]int{1, 2, 2, 3, 4, 4, 3}))
	}
}
