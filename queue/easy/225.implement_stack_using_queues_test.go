package easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/implement-stack-using-queues/description/

// method 1 two queues
// 1) use two queues to implement stack, queue and queueTmp
// 2) queue is used to store the input of push of stack, the front of queue will be the top of the stack
// 3) queueTmp is used to store the result of pop and peek of stack
// 4) when pop or peek, if queueTmp is empty, then move all queue to queueTmp from the front of queue, front of queue will be the top of the queueTmp
// 5) pop or peek the front of queueTmp
// 6) empty is true if queue and queueTmp are empty
// TC = O(N), SC = O(N)
type MyStack struct {
	// * key point: queue is FIFO, we can only get the front of queue, queue[0]
	queue    []int // store the input of push of stack
	queueTmp []int // store the result of pop and peek of stack
}

func ConstructorMyStack() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	/*
		input 1, 2, 3
		1) queue [1], queueTmp [] =>  queue [] , queueTmp [1]
		2) queue [2], queueTmp [1] =>  queue [] , queueTmp [2, 1]
		3) queue [3], queueTmp [2, 1] =>  queue [] , queueTmp [3, 2, 1]
	*/
	this.queue = append(this.queue, x)
	for len(this.queueTmp) > 0 {
		front := this.queueTmp[0]
		this.queueTmp = this.queueTmp[1:]
		this.queue = append(this.queue, front)
	}

	this.queue, this.queueTmp = this.queueTmp, this.queue
}

func (this *MyStack) Pop() int {
	front := this.queueTmp[0]
	this.queueTmp = this.queueTmp[1:]
	return front
}

func (this *MyStack) Top() int {
	return this.queueTmp[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queueTmp) == 0 && len(this.queue) == 0
}

func Test_MyStack_Pop_Top_Empty(t *testing.T) {
	myStack := ConstructorMyStack()

	type args struct {
		x int
	}
	type expected struct {
		resultTop   int
		resultPop   int
		resultEmpty bool
	}
	type testCase struct {
		name           string
		args           args
		myStackMethods func(myStack *MyStack, x int)
		expected       expected
	}

	testCases := []testCase{
		{
			name: "1",
			args: args{
				x: 1,
			},
			myStackMethods: func(myStack *MyStack, x int) {
				myStack.Push(x)
			},
			expected: expected{
				resultTop:   1,
				resultPop:   1,
				resultEmpty: true,
			},
		},
	}

	for _, tc := range testCases {
		tc.myStackMethods(&myStack, tc.args.x)

		assert.Equal(
			t,
			tc.expected.resultTop,
			myStack.Top(),
			fmt.Sprintf("testCase name: %s", tc.name),
		)

		assert.Equal(
			t,
			tc.expected.resultPop,
			myStack.Pop(),
			fmt.Sprintf("testCase name: %s", tc.name),
		)

		assert.Equal(
			t,
			tc.expected.resultEmpty,
			myStack.Empty(),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
