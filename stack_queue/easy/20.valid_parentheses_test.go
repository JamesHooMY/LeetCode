package easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/valid-parentheses/description/

// method 1 stack + map
// 1) use a stack to store the left parentheses
// 2) use a map to store key: right parentheses, value: left parentheses
// 3) if the current character is right parentheses, then use the value (left parentheses) of the map to compare with the top of the stack
// 4) if the value is not equal, then return false. otherwise, pop the top of the stack
// 5) if the current character is left parentheses, then push it into the stack
// 6) finally, if the stack is not empty, then return false. otherwise, return true
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
func isValid1(s string) bool {
	stack := []rune{} // '(', '[', '{'
	charMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, v := range s {
		if val, exist := charMap[v]; exist {
			if len(stack) == 0 || stack[len(stack)-1] != val {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}

	return len(stack) == 0
}

func Test_isValid1(t *testing.T) {
	type args struct {
		s string
	}
	type expected struct {
		result bool
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
				s: "()",
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				s: "()[]{}",
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "3",
			args: args{
				s: "(]",
			},
			expected: expected{
				result: false,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			isValid1(tc.args.s),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
