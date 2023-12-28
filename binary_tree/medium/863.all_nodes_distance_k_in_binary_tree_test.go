package medium

import (
	"fmt"
	"testing"

	"leetcode/binary_tree/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/description/

// method 1 iterative BFS (top-down) + recursive DFS (top-down) Preorder Traversal build adjacency list graph, more easy to understand
// 1) build adjacency list graph
// 2) iterative BFS (top-down) to find all nodes that are k distance from target
// 3) return the result
// TC = O(N), SC = O(N)
func distanceK1[T int](root, target *util.TreeNode[T], k T) []T {
	if root == nil {
		return nil
	}

	// build adjacency list graph
	/*
		example: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, k = 2

		adjacency list graph:
			3: [5, 1]
			5: [3, 6, 2]
			6: [5]
			2: [5, 7, 4]
			7: [2]
			4: [2]
			1: [3, 0, 8]
			0: [1]
			8: [1]
	*/
	listGraph := map[T][]T{} // key: node.Val, value: Val of node's neighbors
	buildListGraph(root, nil, listGraph)

	// iterative BFS (top-down)
	queue := []T{target.Val} // queue stores node.Val
	visited := map[T]bool{target.Val: true} // key: node.Val, value: visited or not
	depth := T(0)
	for len(queue) > 0 {
		levelSize := len(queue)
		if depth == k {
			return queue
		}

		for i := 0; i < levelSize; i++ {
			nodeVal := queue[0]
			queue = queue[1:]
			/*
					depth: 0
					graph:
						5: [3, 6, 2]
					visted:
						5: true
						3: true
						6: true
						2: true
					queue:
						3, 6, 2
				==================================================
					depth: 1
					graph:
						3: [5, 1]
						6: [5]
						2: [5, 7, 4]
					visted:
						5: true
						3: true
						6: true
						2: true
						1: true
						7: true
						4: true
					queue:
						1, 7, 4
				==================================================
					depth: 2
					return queue: [1, 7, 4]
			*/
			for _, neighbor := range listGraph[nodeVal] {
				if !visited[neighbor] {
					queue = append(queue, neighbor)
					visited[neighbor] = true
				}
			}
		}
		depth++
	}

	return []T{}
}

// recursive DFS (top-down) Preorder Traversal build adjacency list graph
func buildListGraph[T int](node *util.TreeNode[T], parent *util.TreeNode[T], listGraph map[T][]T) {
	if node == nil {
		return
	}

	/*
		example: root = [3,5,1,6,2,0,8,null,null,7,4]

		node: 3, parent: nil => graph: {}

		start build graph from child nodes of root node 3:
			node: 5, parent: 3 => graph: {3: [5], 5: [3]}
			node: 6, parent: 5 => graph: {3: [5], 5: [3, 6], 6: [5]}
			node: 2, parent: 5 => graph: {3: [5], 5: [3, 6, 2], 6: [5], 2: [5]}
			node: 7, parent: 2 => graph: {3: [5], 5: [3, 6, 2], 6: [5], 2: [5, 7], 7: [2]}
			node: 4, parent: 2 => graph: {3: [5], 5: [3, 6, 2], 6: [5], 2: [5, 7, 4], 7: [2], 4: [2]}
			node: 1, parent: 3 => graph: {3: [5, 1], 5: [3, 6, 2], 6: [5], 2: [5, 7, 4], 7: [2], 4: [2], 1: [3]}
			node: 0, parent: 1 => graph: {3: [5, 1], 5: [3, 6, 2], 6: [5], 2: [5, 7, 4], 7: [2], 4: [2], 1: [3, 0], 0: [1]}
			node: 8, parent: 1 => graph: {3: [5, 1], 5: [3, 6, 2], 6: [5], 2: [5, 7, 4], 7: [2], 4: [2], 1: [3, 0, 8], 0: [1], 8: [1]}

		final adjacency list graph:
			3: [5, 1]
			5: [3, 6, 2]
			6: [5]
			2: [5, 7, 4]
			7: [2]
			4: [2]
			1: [3, 0, 8]
			0: [1]
			8: [1]
	*/
	if parent != nil {
		listGraph[node.Val] = append(listGraph[node.Val], parent.Val)
		listGraph[parent.Val] = append(listGraph[parent.Val], node.Val)
	}

	buildListGraph(node.Left, node, listGraph)
	buildListGraph(node.Right, node, listGraph)
}

// method 2 recursive DFS (top-down) Preorder Traversal, more logics for memorization
// 1) recursive DFS (top-down) Preorder Traversal to find all nodes that are k distance from target
// 2) return the result
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
func distanceK2[T int](root, target *util.TreeNode[T], k T) []T {
	if root == nil {
		return nil
	}

	result := []T{}
	distanceK2DFS(root, target, k, &result)
	return result
}

func distanceK2DFS[T int](node, target *util.TreeNode[T], k T, result *[]T) T {
	// leaf node with no children
	if node == nil {
		return -1
	}

	// if node is target, find all nodes that are k distance from node, start from node to node.Left and node.Right with DFS Preorder Traversal
	if node == target {
		distanceK2NodeKDFS(node, k, result)
		return 1
	}

	/*
		example: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, k = 2

		node: 3
			1) node.Left: 5, result: []
			2) node == target => distanceK2NodeKDFS(node: 5, k: 2, result: []) => result: [7, 4]
			3) leftDepth(TreeNode 5 <-> root 3): 1

			4) node.Right: 1, result: [7, 4]
			5) leftDepth(TreeNode 0 <-> TreeNode 1): -1, rightDepth(TreeNode 8 <-> TreeNode 1): -1
			6) rightDepth(TreeNode 1 <-> root 3): -1

			7) leftDepth != -1, leftDepth != k
			8) distanceK2NodeKDFS(node: 1, k-leftDepth-1: 0, result: [7, 4]) => result: [7, 4, 1]
			9) return leftDepth + 1: 2
	*/
	leftDepth := distanceK2DFS(node.Left, target, k, result) // distance of (node.Left <-> node)
	rightDepth := distanceK2DFS(node.Right, target, k, result) // distance of (node.Right <-> node)

	// leftDepth == -1 means target is not in the left subtree of node
	if leftDepth != -1 {
		if leftDepth == k {
			*result = append(*result, node.Val)
		} else {
			// find all nodes that are k-leftDepth-1 distance from node.Right
			distanceK2NodeKDFS(node.Right, k-leftDepth-1, result)
		}

		return leftDepth + 1
	}

	// rightDepth == -1 means target is not in the right subtree of node
	if rightDepth != -1 {
		if rightDepth == k {
			*result = append(*result, node.Val)
		} else {
			distanceK2NodeKDFS(node.Left, k-rightDepth-1, result)
		}

		return rightDepth + 1
	}

	return -1
}

func distanceK2NodeKDFS[T int](node *util.TreeNode[T], k T, result *[]T) {
	if node == nil {
		return
	}

	if k == 0 {
		*result = append(*result, node.Val)
		return
	}

	// find all nodes that are k distance from node, start from node.Left and node.Right
	distanceK2NodeKDFS(node.Left, k-1, result)
	distanceK2NodeKDFS(node.Right, k-1, result)
}

func Test_distanceK1(t *testing.T) {
	type args struct {
		root   *util.TreeNode[int]
		target *util.TreeNode[int]
		k      int
	}
	type expected struct {
		result []int
	}
	type testCase struct {
		name     string
		args     args
		expected expected
	}

	roots := []*util.TreeNode[int]{
		util.ArrayToBinaryTree([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}),
		util.ArrayToBinaryTree([]int{1}),
	}

	testCases := []testCase{
		{
			name: "1",
			args: args{
				root:   roots[0],
				target: roots[0].Left,
				k:      2,
			},
			expected: expected{
				result: []int{7, 4, 1},
			},
		},
		{
			name: "2",
			args: args{
				root:   roots[1],
				target: roots[1],
				k:      3,
			},
			expected: expected{
				result: []int{},
			},
		},
	}

	for _, tc := range testCases {
		assert.ElementsMatch(
			t,
			tc.expected.result,
			distanceK1(tc.args.root, tc.args.target, tc.args.k),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_distanceK2(t *testing.T) {
	type args struct {
		root   *util.TreeNode[int]
		target *util.TreeNode[int]
		k      int
	}
	type expected struct {
		result []int
	}
	type testCase struct {
		name     string
		args     args
		expected expected
	}

	roots := []*util.TreeNode[int]{
		util.ArrayToBinaryTree([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}),
		util.ArrayToBinaryTree([]int{1}),
	}

	testCases := []testCase{
		{
			name: "1",
			args: args{
				root:   roots[0],
				target: roots[0].Left,
				k:      2,
			},
			expected: expected{
				result: []int{7, 4, 1},
			},
		},
		{
			name: "2",
			args: args{
				root:   roots[1],
				target: roots[1],
				k:      3,
			},
			expected: expected{
				result: []int{},
			},
		},
	}

	for _, tc := range testCases {
		assert.ElementsMatch(
			t,
			tc.expected.result,
			distanceK2(tc.args.root, tc.args.target, tc.args.k),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_distanceK1(b *testing.B) {
	roots := []*util.TreeNode[int]{
		util.ArrayToBinaryTree([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}),
		util.ArrayToBinaryTree([]int{1}),
	}

	targets := []*util.TreeNode[int]{
		roots[0].Left,
		roots[1],
	}

	ks := []int{
		2,
		3,
	}

	for i := 0; i < b.N; i++ {
		distanceK1(roots[i%2], targets[i%2], ks[i%2])
	}
}

func Benchmark_distanceK2(b *testing.B) {
	roots := []*util.TreeNode[int]{
		util.ArrayToBinaryTree([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}),
		util.ArrayToBinaryTree([]int{1}),
	}

	targets := []*util.TreeNode[int]{
		roots[0].Left,
		roots[1],
	}

	ks := []int{
		2,
		3,
	}

	for i := 0; i < b.N; i++ {
		distanceK2(roots[i%2], targets[i%2], ks[i%2])
	}
}