package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/course-schedule/description/

// method 1 topological sort (BFS, Kahn's algorithm) + hash table + queue
// 1) use a hash table to store the adjacency list of the graph
// 2) use a hash table to store the in-degree of each node
// 3) use a queue to store the nodes with in-degree 0
// 4) use a for loop to scan the prerequisites
// 5) store the graph and in-degree
// 6) use a for loop to scan the nodes with in-degree 0
// 7) store the node in the queue
// 8) use a for loop to scan the queue
// 9) pop the node from the queue
// 10) use a for loop to scan the neighbors of the node
// 11) decrement the in-degree of the neighbor
// 12) if the in-degree of the neighbor is 0, store the neighbor in the queue
// 13) if the queue is empty, return true
// 14) return false
// TC: O(V+E), SC: O(V+E), where V is the number of vertices (nodes), E is the number of edges (prerequisites)
func canFinish1(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)  // key: prerequisite, value: course
	inDegree := make(map[int]int) // key: course, value: in-degree, if value is 0 means can take the course

	// edge[0] is course, edge[1] is prerequisite
	for _, edge := range prerequisites {
		course, pre := edge[0], edge[1]
		graph[pre] = append(graph[pre], course)
		inDegree[course]++
	}

	queue := make([]int, 0) // store the courses with in-degree 0
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	count := 0 // store the number of courses with in-degree 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:] // pop the node from the queue

		count++
		for _, neighbor := range graph[node] {
			inDegree[neighbor]--

			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return count == numCourses
}

// method 2 topological sort (BFS, Kahn's algorithm) + hash table + slice + queue
// 1) use a hash table to store the adjacency list of the graph
// 2) use a slice to store the in-degree of each node
// 3) use a queue to store the nodes with in-degree 0
// 4) use a for loop to scan the prerequisites
// 5) store the graph and in-degree
// 6) use a for loop to scan the nodes with in-degree 0
// 7) store the node in the queue
// 8) use a for loop to scan the queue
// 9) pop the node from the queue
// 10) use a for loop to scan the neighbors of the node
// 11) decrement the in-degree of the neighbor
// 12) if the in-degree of the neighbor is 0, store the neighbor in the queue
// 13) if the queue is empty, return true
// 14) return false
// TC: O(V+E), SC: O(V+E), where V is the number of vertices (nodes), E is the number of edges (prerequisites)
// * this is the best solution for me currently
func canFinish2(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int) // key: prerequisite, value: course
	inDegree := make([]int, numCourses) // index: course, value: in-degree, if value is 0 means can take the course

	// edge[0] is course, edge[1] is prerequisite
	for _, edge := range prerequisites {
		course, pre := edge[0], edge[1]
		graph[pre] = append(graph[pre], course)
		inDegree[course]++
	}

	queue := make([]int, 0) // store the courses with in-degree 0
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	count := 0 // store the number of courses with in-degree 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:] // pop the node from the queue

		count++
		for _, neighbor := range graph[node] {
			inDegree[neighbor]--

			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return count == numCourses
}

func TestCanFinish1(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	type expected struct {
		result bool
	}
	testCases := []struct {
		name     string
		args     args
		expected expected
	}{
		{
			name: "1",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}, {0, 1}},
			},
			expected: expected{
				result: false,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected.result, canFinish1(tc.args.numCourses, tc.args.prerequisites), tc.name)
	}
}

func TestCanFinish2(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	type expected struct {
		result bool
	}
	testCases := []struct {
		name     string
		args     args
		expected expected
	}{
		{
			name: "1",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}, {0, 1}},
			},
			expected: expected{
				result: false,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected.result, canFinish2(tc.args.numCourses, tc.args.prerequisites), tc.name)
	}
}

// benchmark
func BenchmarkCanFinish1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canFinish1(2, [][]int{{1, 0}, {0, 1}})
	}
}

func BenchmarkCanFinish2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canFinish2(2, [][]int{{1, 0}, {0, 1}})
	}
}
