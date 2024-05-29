package util

type GraphNode[T any] struct {
	Val       int
	Neighbors []*GraphNode[T]
}

// adjancent matrix
/*
		a  b  c
	a	0, 1, 1		a -> b, c
	b	1, 0, 1	=>	b -> a, c
	c	1, 1, 0		c -> a, b
*/
// func ArrayToGraph(arr [][]int) []*GraphNode[int] {
// 	nodes := make([]*GraphNode[int], len(arr))
// 	for i := range nodes {
// 		nodes[i] = &GraphNode[int]{Val: i}
// 	}

// 	for i, row := range arr {
// 		for j, v := range row {
// 			if v == 1 {
// 				nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[j])
// 			}
// 		}
// 	}

// 	return nodes
// }
