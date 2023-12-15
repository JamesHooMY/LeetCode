package easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/implement-queue-using-stacks/

// method 1 two stacks
// 1) use two stacks to implement queue, stack and stackTmp
// 2) stack is used to store the input of push of queue, the top of stack will be the end of the queue
// 3) stackTmp is used to store the result of pop and peek of queue
// 4) when pop or peek, if stackTmp is empty, then move all stack to stackTmp from the top of stack, top of stack will be the end of the stackTmp
// 5) pop or peek the top of stackTmp
// 6) empty is true if stack and stackTmp are empty
// TC = O(N), SC = O(N)
type MyQueue struct {
	// * key point: stack is FILO, we can only get the top of stack, stack[len(stack)-1]
	stack    []int // store the input of push of queue
	stackTmp []int // store the result of pop and peek of queue
}

func ConstructorQueue() MyQueue {
	return MyQueue{}
}

// push the new element to the top of stack, the end of the queue
func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)
}

// pop out the first of the queue !!! but stack is FILO
func (this *MyQueue) Pop() int {
	// stack [3, 2, 1], stackTmp [] =>  stack [] , stackTmp [1, 2, 3]
	if len(this.stackTmp) == 0 {
		// move all stack to stackTmp from the top of stack, top of stack will be the end of the stackTmp
		for len(this.stack) > 0 {
			top := this.stack[len(this.stack)-1]
			this.stack = this.stack[:len(this.stack)-1]
			this.stackTmp = append(this.stackTmp, top)
		}
	}

	top := this.stackTmp[len(this.stackTmp)-1]
	this.stackTmp = this.stackTmp[:len(this.stackTmp)-1]

	return top
}

// return the first of the queue !!! but stack is FILO
func (this *MyQueue) Peek() int {
	// stack [3, 2, 1], stackTmp [] =>  stack [] , stackTmp [1, 2, 3]
	if len(this.stackTmp) == 0 {
		// move all stack to stackTmp from the top of stack, top of stack will be the end of the stackTmp
		for len(this.stack) > 0 {
			top := this.stack[len(this.stack)-1]
			this.stack = this.stack[:len(this.stack)-1]
			this.stackTmp = append(this.stackTmp, top)
		}
	}

	return this.stackTmp[len(this.stackTmp)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0 && len(this.stackTmp) == 0
}

func Test_MyQueue_Pop_Peek_Empty(t *testing.T) {
	myQueue := ConstructorQueue()

	type args struct {
		x int
	}
	type expected struct {
		resultPeek  int
		resultPop   int
		resultEmpty bool
	}
	type testCase struct {
		name           string
		args           args
		myQueueMethods func(myQueue *MyQueue, x int)
		expected       expected
	}

	testCases := []testCase{
		{
			name: "1",
			args: args{
				x: 1,
			},
			myQueueMethods: func(myQueue *MyQueue, x int) {
				myQueue.Push(x)
			},
			expected: expected{
				resultPeek:  1,
				resultPop:   1,
				resultEmpty: true,
			},
		},
	}

	for _, tc := range testCases {
		tc.myQueueMethods(&myQueue, tc.args.x)

		assert.Equal(
			t,
			tc.expected.resultPeek,
			myQueue.Peek(),
			fmt.Sprintf("testCase name: %s", tc.name),
		)

		assert.Equal(
			t,
			tc.expected.resultPop,
			myQueue.Pop(),
			fmt.Sprintf("testCase name: %s", tc.name),
		)

		assert.Equal(
			t,
			tc.expected.resultEmpty,
			myQueue.Empty(),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
