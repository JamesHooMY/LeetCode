package util

// singly linked list
type LinkedList struct {
	Head *ListNode
}

// singly linked list node
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

// doubly linked list
type DLinkedList struct {
	Head *DListNode
	Tail *DListNode
}

// doubly linked list node
type DListNode struct {
	Key  int // * this is optional, it is needed to delete the node from the cache in 146._LRU_cache_test.go
	Val  int
	Prev *DListNode
	Next *DListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode // nil
	current := head

	for current != nil {
		next := current.Next // 1) save the next node
		current.Next = prev  // 2) reverse the current node

		// 3) update the prev and current node, for the next iteration
		prev, current = current, next
	}

	return prev
}
