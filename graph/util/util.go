package util

// ListNode for graph
type Node[T any] struct {
	Val       T
	Neighbors []*Node[T]
}

// Adjacency list graph
type GraphList struct {
	Vertices int
	Nodes    []*Node[int]
}

// Adjacency matrix graph
type GraphMatrix struct {
	Vertices int
	Edges    [][]int
}

// ArrayToGraphList constructs a graph from an adjacency list
func ArrayToGraphList(arr [][]int) *GraphList {
	if arr == nil {
		return &GraphList{Vertices: 0}
	}

	vertices := len(arr)
	if vertices == 0 {
		return &GraphList{Vertices: 1, Nodes: []*Node[int]{{Val: 1}}}
	}

	nodes := make([]*Node[int], vertices)
	for i := range nodes {
		nodes[i] = &Node[int]{Val: i + 1}
	}

	for i, neighbors := range arr {
		for _, neighbor := range neighbors {
			nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[neighbor-1])
		}
	}

	return &GraphList{Vertices: vertices, Nodes: nodes}
}
