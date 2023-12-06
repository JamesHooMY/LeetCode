package easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/backspace-string-compare/description/

// method 1
// 1) use stack to store the iterated charactor of string
// 2) if the charactor is '#', pop the last charactor from stack
// 3) compare the stack of two strings
// TC = O(N), SC = O(N)
func backspaceCompare1(s string, t string) bool {
    return processString(s) == processString(t)
}

func processString(str string) string {
	stack := []rune{}

	for _, v := range str {
		if v != '#' {
			stack = append(stack, v)
		} else if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}

	return string(stack)
}

func Test_backspaceCompare1(t *testing.T) {
	type args struct {
		s string
		t string
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
				s: "ab#c",
				t: "ad#c",
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				s: "ab##",
				t: "c#d#",
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "3",
			args: args{
				s: "a##c",
				t: "#a#c",
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "4",
			args: args{
				s: "a#c",
				t: "b",
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
			backspaceCompare1(tc.args.s, tc.args.t),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
