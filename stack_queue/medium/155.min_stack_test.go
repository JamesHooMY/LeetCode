package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/implement-queue-using-stacks/

// method 1 two stacks
// 1) stack is used to store all input
// 2) top of stackMin will keep the min of stack
// 3) when push, if the value is less than or equal to the top of stackMin, then push the value into stackMin
// 4) when pop, if the top of stackMin is equal to the top of stack, then pop the top of stackMin
// 5) getMin will return the top of stackMin
// TC = O(1), SC = O(N)
// * this is the best solution for me currently
type MinStack struct {
	stack    []int // store all input
	stackMin []int // top of stackMin will keep the min of stack
}

func ConstructorMinStack() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	// make sure the top of the stackMin always the min of stack
	if len(this.stackMin) == 0 || val <= this.stackMin[len(this.stackMin)-1] {
		this.stackMin = append(this.stackMin, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}

	top := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	if this.stackMin[len(this.stackMin)-1] == top {
		this.stackMin = this.stackMin[:len(this.stackMin)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return -1 // or any default value
	}

	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.stackMin) == 0 {
		return -1 // or any default value
	}

	return this.stackMin[len(this.stackMin)-1]
}

func Test_MinStack_Pop_Top_GetMin(t *testing.T) {
	minStack := ConstructorMinStack()

	type args struct {
		val int
	}
	type expected struct {
		resultTop    int
		resultGetMin int
	}
	type testCase struct {
		name            string
		args            args
		minStackMethods func(minStack *MinStack, val int)
		expected        expected
	}

	testCases := []testCase{
		{
			name: "1",
			args: args{
				val: 1,
			},
			minStackMethods: func(minStack *MinStack, val int) {
				minStack.Push(val)
			},
			expected: expected{
				resultTop:    1,
				resultGetMin: 1,
			},
		},
		{
			name: "2",
			args: args{
				val: 2,
			},
			minStackMethods: func(minStack *MinStack, val int) {
				minStack.Push(val)
			},
			expected: expected{
				resultTop:    2,
				resultGetMin: 1,
			},
		},
		{
			name: "3",
			args: args{
				val: -2,
			},
			minStackMethods: func(minStack *MinStack, val int) {
				minStack.Push(val)
			},
			expected: expected{
				resultTop:    -2,
				resultGetMin: -2,
			},
		},
		{
			name: "4",
			args: args{},
			minStackMethods: func(minStack *MinStack, val int) {
				minStack.Pop()
			},
			expected: expected{
				resultTop:    2,
				resultGetMin: 1,
			},
		},
	}

	for _, tc := range testCases {
		tc.minStackMethods(&minStack, tc.args.val)

		assert.Equal(
			t,
			tc.expected.resultTop,
			minStack.Top(),
			fmt.Sprintf("testCase name: %s", tc.name),
		)

		assert.Equal(
			t,
			tc.expected.resultGetMin,
			minStack.GetMin(),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
