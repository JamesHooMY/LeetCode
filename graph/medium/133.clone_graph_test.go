package medium

import (
	"testing"

	"leetcode/graph/util"

	"github.com/stretchr/testify/assert"
)

// method 1 recursive DFS (top-down) Preorder Traversal
// use a map to store visited nodes
// clone the node and its neighbors
// if the node is visited, return the clone node
// if the node is not visited, create a new node and mark the original node as visited
// then clone the neighbors
// return the clone node
// TC: O(V+E), SC: O(V+E), where V is the number of vertices (nodes), E is the number of edges (neighbors)
// * this is the best solution for me currently
func cloneGraph1[T int](node *util.Node[T]) *util.Node[T] {
	if node == nil {
		return nil
	}

	visited := make(map[*util.Node[T]]*util.Node[T])

	return dfs1(node, visited)
}

func dfs1[T int](node *util.Node[T], visited map[*util.Node[T]]*util.Node[T]) *util.Node[T] {
	if node == nil {
		return nil
	}

	if _, ok := visited[node]; ok {
		return visited[node]
	}

	clone := &util.Node[T]{Val: node.Val}
	visited[node] = clone // mark original node as visited key, clone as value

	for _, neighbor := range node.Neighbors {
		clone.Neighbors = append(clone.Neighbors, dfs1(neighbor, visited))
	}

	return clone
}

// method 2 iterative DFS (bottom-up) Postorder Traversal, prevent stack overflow for large graphs
// use a map to store visited nodes
// use a stack to store nodes
// clone the node and its neighbors
// if the node is visited, return the clone node
// if the node is not visited, create a new node and mark the original node as visited
// then clone the neighbors
// return the clone node
// TC: O(V+E), SC: O(V+E), where V is the number of vertices (nodes), E is the number of edges (neighbors)
func cloneGraph2[T int](node *util.Node[T]) *util.Node[T] {
	if node == nil {
		return nil
	}

	visited := make(map[*util.Node[T]]*util.Node[T])
	visited[node] = &util.Node[T]{Val: node.Val}

	stack := []*util.Node[T]{node}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, neighbor := range curr.Neighbors {
			if _, ok := visited[neighbor]; !ok {
				visited[neighbor] = &util.Node[T]{Val: neighbor.Val}
				stack = append(stack, neighbor)
			}

			visited[curr].Neighbors = append(visited[curr].Neighbors, visited[neighbor])
		}
	}

	return visited[node]
}

// method 3 BFS (top-down) Level Order Traversal, prevent stack overflow for large graphs
// use a map to store visited nodes
// use a queue to store nodes
// clone the node and its neighbors
// if the node is visited, return the clone node
// if the node is not visited, create a new node and mark the original node as visited
// then clone the neighbors
// return the clone node
// TC: O(V+E), SC: O(V+E), where V is the number of vertices (nodes), E is the number of edges (neighbors)
func cloneGraph3[T int](node *util.Node[T]) *util.Node[T] {
	if node == nil {
		return nil
	}

	visited := make(map[*util.Node[T]]*util.Node[T])
	visited[node] = &util.Node[T]{Val: node.Val}

	queue := []*util.Node[T]{node}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, neighbor := range curr.Neighbors {
			if _, ok := visited[neighbor]; !ok {
				visited[neighbor] = &util.Node[T]{Val: neighbor.Val}
				queue = append(queue, neighbor)
			}

			visited[curr].Neighbors = append(visited[curr].Neighbors, visited[neighbor])
		}
	}

	return visited[node]
}

func Test_cloneGraph1(t *testing.T) {
	type args struct {
		node *util.Node[int]
	}
	type expected struct {
		result *util.Node[int]
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
				node: util.ArrayToGraphList([][]int{
					{2, 4},
					{1, 3},
					{2, 4},
					{1, 3},
				}).Nodes[0],
			},
			expected: expected{
				result: util.ArrayToGraphList([][]int{
					{2, 4},
					{1, 3},
					{2, 4},
					{1, 3},
				}).Nodes[0],
			},
		},
		{
			name: "2",
			args: args{
				node: &util.Node[int]{Val: 1},
			},
			expected: expected{
				result: &util.Node[int]{Val: 1},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := cloneGraph1(tc.args.node)
			assert.Equal(t, tc.expected.result, result)
		})
	}
}

func Test_cloneGraph2(t *testing.T) {
	type args struct {
		node *util.Node[int]
	}
	type expected struct {
		result *util.Node[int]
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
				node: util.ArrayToGraphList([][]int{
					{2, 4},
					{1, 3},
					{2, 4},
					{1, 3},
				}).Nodes[0],
			},
			expected: expected{
				result: util.ArrayToGraphList([][]int{
					{2, 4},
					{1, 3},
					{2, 4},
					{1, 3},
				}).Nodes[0],
			},
		},
		{
			name: "2",
			args: args{
				node: &util.Node[int]{Val: 1},
			},
			expected: expected{
				result: &util.Node[int]{Val: 1},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := cloneGraph2(tc.args.node)
			assert.Equal(t, tc.expected.result, result)
		})
	}
}

func Test_cloneGraph3(t *testing.T) {
	type args struct {
		node *util.Node[int]
	}
	type expected struct {
		result *util.Node[int]
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
				node: util.ArrayToGraphList([][]int{
					{2, 4},
					{1, 3},
					{2, 4},
					{1, 3},
				}).Nodes[0],
			},
			expected: expected{
				result: util.ArrayToGraphList([][]int{
					{2, 4},
					{1, 3},
					{2, 4},
					{1, 3},
				}).Nodes[0],
			},
		},
		{
			name: "2",
			args: args{
				node: &util.Node[int]{Val: 1},
			},
			expected: expected{
				result: &util.Node[int]{Val: 1},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := cloneGraph3(tc.args.node)
			assert.Equal(t, tc.expected.result, result)
		})
	}
}

// benchmark
func Benchmark_cloneGraph1(b *testing.B) {
	node := util.ArrayToGraphList([][]int{
		{2, 4},
		{1, 3},
		{2, 4},
		{1, 3},
	}).Nodes[0]

	for i := 0; i < b.N; i++ {
		cloneGraph1(node)
	}
}

func Benchmark_cloneGraph2(b *testing.B) {
	node := util.ArrayToGraphList([][]int{
		{2, 4},
		{1, 3},
		{2, 4},
		{1, 3},
	}).Nodes[0]

	for i := 0; i < b.N; i++ {
		cloneGraph2(node)
	}
}

func Benchmark_cloneGraph3(b *testing.B) {
	node := util.ArrayToGraphList([][]int{
		{2, 4},
		{1, 3},
		{2, 4},
		{1, 3},
	}).Nodes[0]

	for i := 0; i < b.N; i++ {
		cloneGraph3(node)
	}
}