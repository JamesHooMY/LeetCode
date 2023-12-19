package easy

import (
	"fmt"
	"testing"

	"leetcode/binary_tree/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/same-tree/description/

// method 1 recursive DFS (top-down)
// 1) normally child nodes (Left, Right) of leaf nodes are nil, thus return true directly
// 2) after the above 1) check, if one of p and q is nil, that means exist child nodes (Left, Right) not match, return false
// 3) after the above 1) and 2) check, if both p and q are not nil, check if p.Val == q.Val, if not match, return false
// 4) return isSameTree1(p.Left, q.Left) && isSameTree1(p.Right, q.Right)
// TC = O(N), SC = O(N), N is the height of tree
// * this is the best solution for me currently
func isSameTree1[T int](p *util.TreeNode[T], q *util.TreeNode[T]) bool {
	// 1) normally child nodes (Left, Right) of leaf nodes are nil, thus return true directly
	if p == nil && q == nil {
		return true
	}

	// 2) after the above 1) check, if one of p and q is nil, that means exist child nodes (Left, Right) not match, return false
	if p == nil || q == nil {
		return false
	}

	// 3) after the above 1) and 2) check, if both p and q are not nil, check if p.Val == q.Val, if not match, return false
	if p.Val != q.Val {
		return false
	}

	return isSameTree1(p.Left, q.Left) && isSameTree1(p.Right, q.Right)
}

// method 2 iterative BFS (top-down)
// 1) push p and q into queue
// 2) iterate queue while queue is not empty, pop node from queue, check if p.Val == q.Val
// 3) if p.Val != q.Val, return false
// 4) if p.Left != nil, q.Left != nil, push p.Left and q.Left into queue
// 5) if p.Right != nil, q.Right != nil, push p.Right and q.Right into queue
// 6) return true
// TC = O(N), SC = O(N)
func isSameTree2[T int](p *util.TreeNode[T], q *util.TreeNode[T]) bool {
	queue := []*util.TreeNode[T]{p, q}

	for len(queue) > 0 {
		pNode, qNode := queue[0], queue[1]
		queue = queue[2:] // pop front 2 elements

		if pNode == nil && qNode == nil {
			return true
		}

		if pNode == nil || qNode == nil {
			return false
		}

		if pNode.Val != qNode.Val {
			return false
		}

		queue = append(queue, pNode.Left, qNode.Left, pNode.Right, qNode.Right)
	}

	return true
}

func Test_isSameTree1(t *testing.T) {
	type args struct {
		p *util.TreeNode[int]
		q *util.TreeNode[int]
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
				p: util.ArrayToBinaryTree([]int{1, 2, 3}),
				q: util.ArrayToBinaryTree([]int{1, 2, 3}),
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				p: util.ArrayToBinaryTree([]int{1, 2}),
				q: util.ArrayToBinaryTree([]int{1, -1, 2}),
			},
			expected: expected{
				result: false,
			},
		},
		{
			name: "3",
			args: args{
				p: util.ArrayToBinaryTree([]int{1, 2, 1}),
				q: util.ArrayToBinaryTree([]int{1, 1, 2}),
			},
			expected: expected{
				result: false,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			isSameTree1(tc.args.p, tc.args.q),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_isSameTree2(t *testing.T) {
	type args struct {
		p *util.TreeNode[int]
		q *util.TreeNode[int]
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
				p: util.ArrayToBinaryTree([]int{1, 2, 3}),
				q: util.ArrayToBinaryTree([]int{1, 2, 3}),
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				p: util.ArrayToBinaryTree([]int{1, 2}),
				q: util.ArrayToBinaryTree([]int{1, -1, 2}),
			},
			expected: expected{
				result: false,
			},
		},
		{
			name: "3",
			args: args{
				p: util.ArrayToBinaryTree([]int{1, 2, 1}),
				q: util.ArrayToBinaryTree([]int{1, 1, 2}),
			},
			expected: expected{
				result: false,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			isSameTree2(tc.args.p, tc.args.q),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_isSameTree1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isSameTree1(util.ArrayToBinaryTree([]int{1, 2, 3}), util.ArrayToBinaryTree([]int{1, 2, 3}))
	}
}

func Benchmark_isSameTree2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isSameTree2(util.ArrayToBinaryTree([]int{1, 2, 3}), util.ArrayToBinaryTree([]int{1, 2, 3}))
	}
}
