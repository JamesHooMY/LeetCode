package hard

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-frequency-stack/

// method 1
type FreqStack struct {
	valFreqMap      map[int]int   // key: val, value: freq
	freqStackValMap map[int][]int // key: freq, value: stack of val
	maxFreq         int           // keep the max freq
}

func ConstructorFreqStack() FreqStack {
	return FreqStack{
		valFreqMap:      map[int]int{},
		freqStackValMap: map[int][]int{},
		maxFreq:         0,
	}
}

func (this *FreqStack) Push(val int) {
	this.valFreqMap[val]++
	freq := this.valFreqMap[val]
	if freq > this.maxFreq {
		this.maxFreq = freq
	}
	this.freqStackValMap[freq] = append(this.freqStackValMap[freq], val)
}

func (this *FreqStack) Pop() int {
	if len(this.freqStackValMap) == 0 {
		return -1
	}

	topVal := this.freqStackValMap[this.maxFreq][len(this.freqStackValMap[this.maxFreq])-1]                                // get the top of stack
	this.freqStackValMap[this.maxFreq] = this.freqStackValMap[this.maxFreq][:len(this.freqStackValMap[this.maxFreq])-1] // pop the top of stack

	if len(this.freqStackValMap[this.maxFreq]) == 0 {
		delete(this.freqStackValMap, this.maxFreq)
		this.maxFreq--
	}

	this.valFreqMap[topVal]--

	return topVal
}

func Test_FreqStack_Pop(t *testing.T) {
	freqStack := ConstructorFreqStack()

	type args struct {
		val int
	}
	type expected struct {
		resultPop int
	}
	type testCase struct {
		name             string
		args             args
		freqStackMethods func(freqStack *FreqStack, val int)
		expected         expected
	}

	testCases := []testCase{
		{
			name: "1",
			args: args{
				val: 1,
			},
			freqStackMethods: func(freqStack *FreqStack, val int) {
				freqStack.Push(val)
			},
			expected: expected{
				resultPop: 1,
			},
		},
		{
			name: "2",
			args: args{
				val: 2,
			},
			freqStackMethods: func(freqStack *FreqStack, val int) {
				freqStack.Push(val)
			},
			expected: expected{
				resultPop: 2,
			},
		},
	}

	for _, tc := range testCases {
		tc.freqStackMethods(&freqStack, tc.args.val)

		assert.Equal(
			t,
			tc.expected.resultPop,
			freqStack.Pop(),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
