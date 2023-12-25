package medium

import (
	"fmt"
	"testing"

	"leetcode/binary_tree/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

// method 1 iterative BFS (top-down)
// 1) push root into queue
// 2) iterate queue while queue is not empty
// 3) iterate levelSize times, pop node from queue, push left and right of node into queue
// 4) append node.Val into levelVals, if level is odd, reverse the order of levelVals from right to left
// 4) append levelVals into result, level++
// 5) return result
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
func zigzagLevelOrder1[T int](root *util.TreeNode[T]) [][]T {
	result := [][]T{}
	if root == nil {
		return result
	}

	queue := []*util.TreeNode[T]{root}
	level := 0
	for len(queue) > 0 {
		levelSize := len(queue)
		levelVals := make([]T, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			idx := i
			// if level is odd, reverse the order of levelVals from right to left
			if level%2 == 1 {
				idx = levelSize - 1 - i
			}

			levelVals[idx] = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}

		result = append(result, levelVals)
		level++
	}

	return result
}

func Test_zigzagLevelOrder1(t *testing.T) {
	type args struct {
		root *util.TreeNode[int]
	}
	type expected struct {
		result [][]int
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
				result: [][]int{{3}, {20, 9}, {15, 7}},
			},
		},
		{
			name: "2",
			args: args{
				root: util.ArrayToBinaryTree([]int{1}),
			},
			expected: expected{
				result: [][]int{{1}},
			},
		},
		{
			name: "3",
			args: args{
				root: util.ArrayToBinaryTree([]int{}),
			},
			expected: expected{
				result: [][]int{},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			zigzagLevelOrder1(tc.args.root),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
