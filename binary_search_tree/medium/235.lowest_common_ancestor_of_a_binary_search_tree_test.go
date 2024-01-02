package medium

import (
	"fmt"
	"testing"

	"leetcode/binary_tree/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/

// method 1 recursive DFS (top-down) Postorder Traversal with BST feature
// 1) if root is nil, return nil
// 2) if root is p or q, return root
// 3) if root.Val > p.Val && root.Val > q.Val, call lowestCommonAncestor3(root.Left, p, q)
// 4) if root.Val < p.Val && root.Val < q.Val, call lowestCommonAncestor3(root.Right, p, q)
// 5) return root
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
func lowestCommonAncestor1[T int](root, p, q *util.TreeNode[T]) *util.TreeNode[T] {
	if root == nil {
		return nil
	}

	// if Val of p and q are both smaller than Val of root, LCA must be in left subtree
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor1(root.Left, p, q)
	}

	// if Val of p and q are both larger than Val of root, LCA must be in right subtree
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor1(root.Right, p, q)
	}

	/*
		1) if Val of p or q is equal to Val of root, LCA must be root
		2) p.Val < root.Val < q.Val || q.Val < root.Val < p.Val, LCA must be root
	*/
	return root
}

func Test_lowestCommonAncestor1(t *testing.T) {
	type args struct {
		root *util.TreeNode[int]
		p    *util.TreeNode[int]
		q    *util.TreeNode[int]
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
		util.ArrayToBinaryTree([]int{6,2,8,0,4,7,9,-1,-1,3,5}),
		util.ArrayToBinaryTree([]int{6,2,8,0,4,7,9,-1,-1,3,5}),
		util.ArrayToBinaryTree([]int{2,1}),
	}

	testCases := []testCase{
		{
			name: "1",
			args: args{
				root: roots[0],
				p:    roots[0].Left,
				q:    roots[0].Right,
			},
			expected: expected{
				result: roots[0],
			},
		},
		{
			name: "2",
			args: args{
				root: roots[1],
				p:    roots[1].Left,
				q:    roots[1].Left.Right,
			},
			expected: expected{
				result: roots[1].Left,
			},
		},
		{
			name: "3",
			args: args{
				root: roots[2],
				p:    roots[2],
				q:    roots[2].Left,
			},
			expected: expected{
				result: roots[2],
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			lowestCommonAncestor1(tc.args.root, tc.args.p, tc.args.q),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
