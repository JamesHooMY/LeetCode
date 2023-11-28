package medium

import (
	"testing"

	"leetcode/linked_list/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/lru-cache/description/

// method 1 double linked list + hash table
// 1) use a hash table to store the key and the node
// 2) use a double linked list to store the nodes
// 3) when get a node, move the node to the head of the double linked list
// 4) when put a node, if the node is already in the hash table, then update the value of the node and move the node to the head of the double linked list
// 5) when put a node, if the node is not in the hash table, then add the node to the head of the double linked list, and add the key and the node to the hash table
// 6) when put a node, if the size of the hash table is larger than the capacity, then remove the tail node from the double linked list and the hash table
// TC = O(1), SC = O(N)
// * this is the best solution for me currently
type LRUCache struct {
	capacity int                     // the capacity of the cache
	cache    map[int]*util.DListNode // key: key, value: node
	head     *util.DListNode         // dummy head
	tail     *util.DListNode         // dummy tail
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*util.DListNode),
		head:     &util.DListNode{}, // dummy head
		tail:     &util.DListNode{}, // dummy tail
	}

	// double linked list, head(dummy head) <-> tail(dummy tail)
	l.head.Next = l.tail // head.Prev = nil
	l.tail.Prev = l.head // tail.Next = nil
	return l
}

// * keep the most recently used node at the head
func (this *LRUCache) Get(key int) int {
	if dListNode, ok := this.cache[key]; ok {
		// move the dListNode to the head
		this.moveToHead(dListNode)
		return dListNode.Val
	}

	return -1
}

// * remove the least recently used node from the tail
func (this *LRUCache) Put(key int, value int) {
	if dListNode, ok := this.cache[key]; ok {
		// update the value
		dListNode.Val = value
		// move the dListNode to the head
		this.moveToHead(dListNode)
	} else {
		dListNode := &util.DListNode{
			Key: key,
			Val: value,
		}
		this.cache[key] = dListNode
		this.addNodeToHead(dListNode)

		if len(this.cache) > this.capacity {
			// remove the least recently used node from the tail
			tailNode := this.removeNodeFromTail()
			delete(this.cache, tailNode.Key)
		}
	}
}

func (this *LRUCache) moveToHead(dListNode *util.DListNode) {
	this.removeNode(dListNode)
	this.addNodeToHead(dListNode)
}

func (this *LRUCache) removeNode(dListNode *util.DListNode) {
	// nil <- 1(dummy head) <-> 3 <-> 2(dummy tail) -> nil
	// remove 3
	// nil <- 1(dummy head) <-> 2(dummy tail) -> nil
	dListNode.Prev.Next = dListNode.Next // update the next of the prev node of the dListNode
	dListNode.Next.Prev = dListNode.Prev // update the prev of the next node of the dListNode
}

// * always add the new node at next of the dummy head
func (this *LRUCache) addNodeToHead(dListNode *util.DListNode) {
	// nil <- 1(dummy head) <-> 2(dummy tail) -> nil
	// add 3
	// 1(dummy head) <- 3(dListNode) -> 2(dummy tail); nil <- 1(dummy head) <-> 2(dummy tail) -> nil
	dListNode.Prev = this.head      // update the prev of the new node to the dummy head
	dListNode.Next = this.head.Next // update the next of the new node to the next of the dummy head

	// nil <- 1(dummy head) <-> 3(dListNode) <-> 2(dummy tail) -> nil
	this.head.Next.Prev = dListNode // update the prev of the next node of the dummy head
	this.head.Next = dListNode      // update the next of the dummy head
}

// * always remove the node from prev of the dummy tail
func (this *LRUCache) removeNodeFromTail() *util.DListNode {
	// nil <- 1(dummy head) <-> 3(tailNode) <-> 2(dummy tail) -> nil
	// remove 3
	// nil <- 1(dummy head) <-> 2(dummy tail) -> nil
	tailNode := this.tail.Prev
	this.removeNode(tailNode)
	return tailNode
}

// * test cases are sequence dependent
func Test_LRUCache_Get(t *testing.T) {
	lruCache := Constructor(2)

	type args struct {
		key   int
		value int
	}
	type expected struct {
		result int
	}
	type testCase struct {
		name            string
		args            args
		lruCacheMethods func(lruCache *LRUCache, key, value int)
		expected        expected
	}

	testCases := []testCase{
		{
			name: "1",
			args: args{
				key: 1,
			},
			lruCacheMethods: func(lruCache *LRUCache, key, value int) {},
			expected: expected{
				result: -1,
			},
		},
		{
			name: "2",
			args: args{
				key:   1,
				value: 1,
			},
			lruCacheMethods: func(lruCache *LRUCache, key, value int) {
				lruCache.Put(key, value)
			},
			expected: expected{
				result: 1,
			},
		},
		{
			name: "3",
			args: args{
				key:   2,
				value: 2,
			},
			lruCacheMethods: func(lruCache *LRUCache, key, value int) {
				lruCache.Put(key, value)
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "4",
			args: args{
				key:   2,
				value: 0,
			},
			lruCacheMethods: func(lruCache *LRUCache, key, value int) {
				lruCache.Put(key, value)
			},
			expected: expected{
				result: 0,
			},
		},
		{
			name: "5",
			args: args{
				key:   3,
				value: 3,
			},
			lruCacheMethods: func(lruCache *LRUCache, key, value int) {
				lruCache.Put(key, value)
			},
			expected: expected{
				result: 3,
			},
		},
		{
			name: "6",
			args: args{
				key: 1,
			},
			lruCacheMethods: func(lruCache *LRUCache, key, value int) {},
			expected: expected{
				result: -1,
			},
		},
		{
			name: "7",
			args: args{
				key: 2,
			},
			lruCacheMethods: func(lruCache *LRUCache, key, value int) {},
			expected: expected{
				result: 0,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.lruCacheMethods(&lruCache, tc.args.key, tc.args.value)
			result := lruCache.Get(tc.args.key)
			assert.Equal(t, tc.expected.result, result)
		})
	}
}
