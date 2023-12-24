package medium

import (
	"fmt"
	"testing"

	"leetcode/binary_tree/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/

// method 1 recursive DFS (bottom-up) Postorder Traversal
// 1) if root is nil, return nil
// 2) if root is p or q, return root
// 3) call lowestCommonAncestor1(root.Left, p, q) and assign to left
// 4) call lowestCommonAncestor1(root.Right, p, q) and assign to right
// 5) if left is not nil and right is not nil, return root
// 6) if left is nil, return right
// TC = O(N), SC = O(N), N is the height of tree
// * this is the best solution for me currently
func lowestCommonAncestor1[T int](root, p, q *util.TreeNode[T]) *util.TreeNode[T] {
	if root == nil {
		return nil
	}

	// * this root must be LCA of p and q
	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor1(root.Left, p, q)
	right := lowestCommonAncestor1(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	// * right might be nil or not nil
	return right
}

// method 2 iterative BFS (top-down)
// 1) push root into queue, and add root into parent map with nil as value
// 2) iterate queue while queue is not empty
// 3) iterate levelSize times, pop node from queue, push left and right of node into queue, and add left and right into parent map with node as value
// 4) add p and its ancestors into ancestors map
// 5) iterate q and its ancestors, if q is in ancestors map, return q
// 6) return nil
// TC = O(N), SC = O(N)
func lowestCommonAncestor2[T int](root, p, q *util.TreeNode[T]) *util.TreeNode[T] {
	if root == nil {
		return nil
	}

	queue := []*util.TreeNode[T]{root}

	// * this map is used to store parent of each node
	parent := make(map[*util.TreeNode[T]]*util.TreeNode[T])
	// root has no parent
	parent[root] = nil

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil {
			parent[node.Left] = node
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			parent[node.Right] = node
			queue = append(queue, node.Right)
		}
	}

	// * this map is used to store all ancestors of p
	ancestors := make(map[*util.TreeNode[T]]bool)
	for p != nil {
		ancestors[p] = true

		// update p to its parent
		p = parent[p]
	}

	// * find the first ancestor of q which is also ancestor of p
	for q != nil {
		if ancestors[q] {
			return q
		}

		// update q to its parent
		q = parent[q]
	}

	return nil
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
		util.ArrayToBinaryTree([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}),
		util.ArrayToBinaryTree([]int{1, 2}),
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
				p:    roots[1],
				q:    roots[1].Left,
			},
			expected: expected{
				result: roots[1],
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

func Test_lowestCommonAncestor2(t *testing.T) {
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
		util.ArrayToBinaryTree([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}),
		util.ArrayToBinaryTree([]int{1, 2}),
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
				p:    roots[1],
				q:    roots[1].Left,
			},
			expected: expected{
				result: roots[1],
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			lowestCommonAncestor2(tc.args.root, tc.args.p, tc.args.q),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_lowestCommonAncestor1(b *testing.B) {
	roots := []*util.TreeNode[int]{
		util.ArrayToBinaryTree([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}),
		util.ArrayToBinaryTree([]int{1, 2}),
	}

	for i := 0; i < b.N; i++ {
		lowestCommonAncestor1(roots[0], roots[0].Left, roots[0].Right)
	}
}

func Benchmark_lowestCommonAncestor2(b *testing.B) {
	roots := []*util.TreeNode[int]{
		util.ArrayToBinaryTree([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}),
		util.ArrayToBinaryTree([]int{1, 2}),
	}

	for i := 0; i < b.N; i++ {
		lowestCommonAncestor2(roots[0], roots[0].Left, roots[0].Right)
	}
}
