package medium

import (
	"fmt"
	"testing"

	"leetcode/binary_tree/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/path-sum-iii/description/

// method 1 recursive DFS (top-down) Preorder Traversal
// 1) if root is nil, return 0
// 2) call pathSumIII1DFS(root, targetSum) + pathSumIII1(root.Left, targetSum) + pathSumIII1(root.Right, targetSum)
// 3) return result
// TC = O(N^2), SC = O(N)
func pathSumIII1[T int](root *util.TreeNode[T], targetSum T) T {
	if root == nil {
		return 0
	}

	// pathSumIII1(root.Left, targetSum) + pathSumIII1(root.Right, targetSum) will cause TC = O(N^2)
	return pathSumIII1DFS(root, targetSum) + pathSumIII1(root.Left, targetSum) + pathSumIII1(root.Right, targetSum)
}

func pathSumIII1DFS[T int](node *util.TreeNode[T], sum T) T {
	if node == nil {
		return 0
	}

	// count will be initial with 0 at every pathSumIII1DFS call, due to every may be a path
	var count T
	if node.Val == sum {
		count++
	}

	count += pathSumIII1DFS(node.Left, sum-node.Val) + pathSumIII1DFS(node.Right, sum-node.Val)

	return count
}

// method 2 recursive DFS (top-down) Preorder Traversal + prefix sum hash map, prefix sum like 560.subarray_sum_equals_k.go
// 1) if root is nil, return 0
// 2) define a map to store prefixSum and count, key: prefixSum, value: count
// 3) call pathSumIII2DFS(root, targetSum, 0, curSumMap)
// 4) return result
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
func pathSumIII2[T int](root *util.TreeNode[T], targetSum T) T {
	if root == nil {
		return 0
	}

	sumCountMap := map[T]int{0: 1} // key: curSum, value: count
	return pathSumIII2DFS(root, targetSum, 0, sumCountMap)
}

func pathSumIII2DFS[T int](node *util.TreeNode[T], targetSum T, curSum T, sumCountMap map[T]int) T {
	if node == nil {
		return 0
	}

	curSum += node.Val
	var count T
	if sumCount, ok := sumCountMap[curSum-targetSum]; ok {
		count = T(sumCount)
	}

	// * this is the key point, add curSum for next node traversal
	sumCountMap[curSum]++

	count += pathSumIII2DFS(node.Left, targetSum, curSum, sumCountMap)
	count += pathSumIII2DFS(node.Right, targetSum, curSum, sumCountMap)

	// * this is the key point, remove curSum from map for backtracking, because curSum is only valid in current path. After left subtree traversal, curSum will be back to its original value, then continue to traverse right subtree
	sumCountMap[curSum]--

	return count
}

func Test_pathSumIII1(t *testing.T) {
	type args struct {
		root *util.TreeNode[int]
		sum  int
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
				root: util.ArrayToBinaryTree([]int{10, 5, -3, 3, 2, -1, 11, 3, -2, -1, 1}),
				sum:  8,
			},
			expected: expected{
				result: 3,
			},
		},
		{
			name: "2",
			args: args{
				root: util.ArrayToBinaryTree([]int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, 5, 1}),
				sum:  22,
			},
			expected: expected{
				result: 3,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			pathSumIII1(tc.args.root, tc.args.sum),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_pathSumIII2(t *testing.T) {
	type args struct {
		root *util.TreeNode[int]
		sum  int
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
				root: util.ArrayToBinaryTree([]int{10, 5, -3, 3, 2, -1, 11, 3, -2, -1, 1}),
				sum:  8,
			},
			expected: expected{
				result: 3,
			},
		},
		{
			name: "2",
			args: args{
				root: util.ArrayToBinaryTree([]int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, 5, 1}),
				sum:  22,
			},
			expected: expected{
				result: 3,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			pathSumIII2(tc.args.root, tc.args.sum),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
