package easy

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/evaluate-reverse-polish-notation/

// method 1 stack
// 1) use a stack to store the number
// 2) if the current character is operator, then pop the top two numbers from stack, and calculate the result
// 3) push the result into stack
// 4) finally, the top of stack is the result
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
func evalRPN1(tokens []string) int {
	stack := []int{}

	for _, v := range tokens {
		switch v {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return 0
			}

			second := stack[len(stack)-1]
			first := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch v {
			case "+":
				stack = append(stack, first+second)
			case "-":
				stack = append(stack, first-second)
			case "*":
				stack = append(stack, first*second)
			case "/":
				stack = append(stack, first/second)
			}
		default:
			val, _ := strconv.Atoi(v)
			stack = append(stack, val)
		}
	}

	if len(stack) != 1 {
		return 0
	}

	return stack[0]
}

func Test_evalRPN1(t *testing.T) {
	type args struct {
		tokens []string
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
				tokens: []string{"2", "1", "+", "3", "*"},
			},
			expected: expected{
				result: 9,
			},
		},
		{
			name: "2",
			args: args{
				tokens: []string{"4", "13", "5", "/", "+"},
			},
			expected: expected{
				result: 6,
			},
		},
		{
			name: "3",
			args: args{
				tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			},
			expected: expected{
				result: 22,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			evalRPN1(tc.args.tokens),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
