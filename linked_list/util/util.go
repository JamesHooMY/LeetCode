package util

import (
	"math/rand"
	"time"
)

type LinkedList struct {
	Head *ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func ArrayToSinglyLinkedList(arr []int) *LinkedList {
	ll := &LinkedList{}
	current := ll.Head

	for _, value := range arr {
		if current == nil {
			ll.Head = &ListNode{Val: value, Next: nil}
			current = ll.Head
		} else {
			current.Next = &ListNode{Val: value, Next: nil}
			current = current.Next
		}
	}

	return ll
}

func ArrayToCycleLinkedList(arr []int) *LinkedList {
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

	// Randomly select a node as the start of the cycle
	seed := time.Now().UnixNano()
	randomGenerator := rand.New(rand.NewSource(seed))
	cycleStartNode := nodes[randomGenerator.Intn(len(arr))]

	// Connect the last node to the selected start node to create the random cycle
	current.Next = cycleStartNode

	return &LinkedList{Head: head}
}
