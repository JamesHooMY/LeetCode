package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/basic-calculator-ii/description/

// method 1 stack
// 1) use a stack to store the iterated number
// 2) if the current character is number, then calculate the current number
// 3) if the current character is operator, then calculate the result with the current number and the top of stack number by the current operator
// 4) after calculation, push the result into stack, and reset the current number and update the current operator for next calculation
// 5) finally, the sum of stack is the result
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
func calculate1(s string) int {
	stack := []int{} // store the iterated number
	curNum := 0      // current number
	curSign := '+'   // curSign for current calculation

	for i, char := range s {
		isDigit := char >= '0' && char <= '9'

		if isDigit {
			curNum = curNum*10 + int(char-'0')
		}

		// need to check the char is not ' ' or the last charactor
		if char != ' ' && !isDigit || i == len(s)-1 {
			// calculate the current number and store it into stack
			switch curSign {
			case '+':
				stack = append(stack, curNum)
			case '-':
				stack = append(stack, -curNum)
			case '*':
				stack[len(stack)-1] *= curNum
			case '/':
				stack[len(stack)-1] /= curNum
			}

			curSign = char // update the current curSign for next calculation
			curNum = 0     // reset num
		}
	}

	// calculate the result
	result := 0
	for _, num := range stack {
		result += num
	}

	return result
}

func Test_calculate1(t *testing.T) {
	type args struct {
		s string
	}
	type expected struct {
		result int
	}
	type testCase struct {
		name     string
		args     args
		expected expected
	}

	testCases := []testCase{
		{
			name: "1",
			args: args{
				s: "3+2*2",
			},
			expected: expected{
				result: 7,
			},
		},
		{
			name: "2",
			args: args{
				s: " 3/2 ",
			},
			expected: expected{
				result: 1,
			},
		},
		{
			name: "3",
			args: args{
				s: " 3+5 / 2 ",
			},
			expected: expected{
				result: 5,
			},
		},
		{
			name: "4",
			args: args{
				s: "42",
			},
			expected: expected{
				result: 42,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			calculate1(tc.args.s),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
