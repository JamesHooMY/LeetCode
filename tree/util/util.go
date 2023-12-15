package util

type TreeNode[T any] struct {
    Val   T
    Left  *TreeNode[T]
    Right *TreeNode[T]
}

/*
	sequencial storage to binary tree
	1) suitable for perfect binary tree
	2) waste space to store nil node in non-perfect binary tree (eg. full binary tree and complete binary tree)
*/
func ArrayToBinaryTree[T int](arr []T) *TreeNode[T] {
	if len(arr) == 0 {
		return nil
	}

	root := &TreeNode[T]{Val: arr[0]}
	queue := []*TreeNode[T]{root}

	for i := 1; i < len(arr); i++ {
		node := queue[0]
		queue = queue[1:]

		if arr[i] != -1 {
			node.Left = &TreeNode[T]{Val: arr[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(arr) && arr[i] != -1 {
			node.Right = &TreeNode[T]{Val: arr[i]}
			queue = append(queue, node.Right)
		}
	}

	return root
}