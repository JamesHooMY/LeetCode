package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-height-trees/description/

// method 1 Topological Sort (BFS, Layer by Layer Removal) + Hash Table + Queue
// 1) if n == 1, return []int{0}
// 2) use a hash table to store the adjacency list of the graph
// 3) use a hash table to store the in-degree of each node
// 4) use a queue to store the nodes with in-degree 1
// 5) use a for loop to scan the edges
// 6) store the graph and in-degree
// 7) use a for loop to scan the nodes with in-degree 1
// 8) store the node in the queue
// 9) use a while loop to scan the queue
// 10) if n <= 2, return the queue
// 11) use a for loop to scan the queue
// 12) pop the node from the queue
// 13) use a for loop to scan the neighbors of the node
// 14) decrement the in-degree of the neighbor
// 15) if the in-degree of the neighbor is 1, store the neighbor in the queue
// 16) return the queue
// TC: O(N), SC: O(N), where N is the number of nodes
func findMinHeightTrees1(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	// adjacency list
	graph := map[int][]int{}  // key: node, value: neighbors
	inDegree := map[int]int{} // key: node, value: in-degree

	// build the graph and in-degree
	for _, edge := range edges {
		n1, n2 := edge[0], edge[1]
		graph[n1] = append(graph[n1], n2)
		graph[n2] = append(graph[n2], n1)
		inDegree[n1]++
		inDegree[n2]++
	}

	queue := []int{} // store the nodes with in-degree 1
	for node, degree := range inDegree {
		if degree == 1 {
			queue = append(queue, node)
		}
	}

	// layer by layer removal
	for len(queue) > 0 {
		size := len(queue)
		if n <= 2 { // * if n <= 2, the remaining nodes are the roots
			return queue
		}

		n -= size // * remove the nodes with in-degree 1

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			for _, neighbor := range graph[node] {
				inDegree[neighbor]--
				if inDegree[neighbor] == 1 {
					queue = append(queue, neighbor)
				}
			}
		}

	}

	return queue
}

// method 2 Topological Sort (BFS, Layer by Layer Removal) + Slice + Queue
// 1) if n == 1, return []int{0}
// 2) use a slice to store the adjacency list of the graph
// 3) use a slice to store the in-degree of each node
// 4) use a queue to store the nodes with in-degree 1
// 5) use a for loop to scan the edges
// 6) store the graph and in-degree
// 7) use a for loop to scan the nodes with in-degree 1
// 8) store the node in the queue
// 9) use a while loop to scan the queue
// 10) if n <= 2, return the queue
// 11) use a for loop to scan the queue
// 12) pop the node from the queue
// 13) use a for loop to scan the neighbors of the node
// 14) decrement the in-degree of the neighbor
// 15) if the in-degree of the neighbor is 1, store the neighbor in the queue
// 16) return the queue
// TC: O(N), SC: O(N), where N is the number of nodes
// * this is the best solution for me currently
func findMinHeightTrees2(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	// adjacency list
	graph := make([][]int, n)  // key: node, value: neighbors
	inDegree := make([]int, n) // key: node, value: in-degree

	// build the graph and in-degree
	for _, edge := range edges {
		n1, n2 := edge[0], edge[1]
		graph[n1] = append(graph[n1], n2)
		graph[n2] = append(graph[n2], n1)
		inDegree[n1]++
		inDegree[n2]++
	}

	queue := []int{} // store the nodes with in-degree 1
	for node, degree := range inDegree {
		if degree == 1 {
			queue = append(queue, node)
		}
	}

	// layer by layer removal
	for len(queue) > 0 {
		size := len(queue)
		if n <= 2 { // * if n <= 2, the remaining nodes are the roots
			return queue
		}

		n -= size // * remove the nodes with in-degree 1

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			for _, neighbor := range graph[node] {
				inDegree[neighbor]--
				if inDegree[neighbor] == 1 {
					queue = append(queue, neighbor)
				}
			}
		}

	}

	return queue
}

func Test_findMinHeightTrees1(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	type expected struct {
		result []int
	}
	testCases := []struct {
		name     string
		args     args
		expected expected
	}{
		{
			name: "1",
			args: args{
				n:     4,
				edges: [][]int{{1, 0}, {1, 2}, {1, 3}},
			},
			expected: expected{
				result: []int{1},
			},
		},
		{
			name: "2",
			args: args{
				n:     6,
				edges: [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}},
			},
			expected: expected{
				result: []int{3, 4},
			},
		},
		{
			name: "3",
			args: args{
				n:     1,
				edges: [][]int{},
			},
			expected: expected{
				result: []int{0},
			},
		},
	}

	for _, tc := range testCases {
		assert.ElementsMatch(t, tc.expected.result, findMinHeightTrees1(tc.args.n, tc.args.edges), tc.name)
	}
}

func Test_findMinHeightTrees2(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	type expected struct {
		result []int
	}
	testCases := []struct {
		name     string
		args     args
		expected expected
	}{
		{
			name: "1",
			args: args{
				n:     4,
				edges: [][]int{{1, 0}, {1, 2}, {1, 3}},
			},
			expected: expected{
				result: []int{1},
			},
		},
		{
			name: "2",
			args: args{
				n:     6,
				edges: [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}},
			},
			expected: expected{
				result: []int{3, 4},
			},
		},
		{
			name: "3",
			args: args{
				n:     1,
				edges: [][]int{},
			},
			expected: expected{
				result: []int{0},
			},
		},
	}

	for _, tc := range testCases {
		assert.ElementsMatch(t, tc.expected.result, findMinHeightTrees2(tc.args.n, tc.args.edges), tc.name)
	}
}

// benchmark
func Benchmark_findMinHeightTrees1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findMinHeightTrees1(4, [][]int{{1, 0}, {1, 2}, {1, 3}})
	}
}

func Benchmark_findMinHeightTrees2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findMinHeightTrees2(4, [][]int{{1, 0}, {1, 2}, {1, 3}})
	}
}
