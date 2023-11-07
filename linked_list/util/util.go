package util

type LinkedList struct {
	Head *ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// pos = -1 means no cycle, otherwise pos is the index of the node where the cycle begins
func ArrayToCycleOrSinglyLinkedList(arr []int, pos int) *LinkedList {
	if len(arr) == 0 {
		return &LinkedList{}
	}

	// Create the first node
	head := &ListNode{Val: arr[0], Next: nil}
	current := head

	// Create the rest of the nodes and add them to the linked list
	nodes := []*ListNode{head}

	for i := 1; i < len(arr); i++ {
		node := &ListNode{Val: arr[i], Next: nil}
		current.Next = node
		current = node
		nodes = append(nodes, node)
	}

	// If pos is -1, then return the linked list without cycle
	if pos < 0 {
		return &LinkedList{Head: head}
	}

	// Decide the start node of the cycle
	cycleStartNode := nodes[pos]

	// Connect the last node to the selected start node to create the cycle
	current.Next = cycleStartNode

	return &LinkedList{Head: head}
}
