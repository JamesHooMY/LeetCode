package easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/backspace-string-compare/description/

// method 1 stack, this is more easy to understand
// 1) use stackChar to store the iterated charactor of string
// 2) if the charactor is '#', pop the last charactor from stackChar
// 3) compare the stackChar of two strings
// TC = O(N), SC = O(N)
func backspaceCompare1(s string, t string) bool {
    return backSpaceProcess(s) == backSpaceProcess(t)
}

func backSpaceProcess(str string) string {
	stackChar := []rune{}

	for _, char := range str {
		if char != '#' {
			stackChar = append(stackChar, char)
		} else if len(stackChar) > 0 {
			stackChar = stackChar[:len(stackChar)-1]
		}
	}

	return string(stackChar)
}

// method 2 two pointers
// 1) use two pointers to compare the string from the end to the start
// 2) if the charactor is '#', skip the next charactor
// 3) compare the two charactors
// TC = O(N), SC = O(1)
// * this is the best solution for me currently
func backspaceCompare2(s string, t string) bool {
	sIdx, tIdx := len(s)-1, len(t)-1

	for sIdx >= 0 || tIdx >= 0 {
		sIdx = backSpaceNextChar(s, sIdx)
		tIdx = backSpaceNextChar(t, tIdx)

		// '' == '' is true
		if sIdx < 0 && tIdx < 0 {
			return true
		}

		// if one of the string is empty, return false
		if sIdx < 0 || tIdx < 0 || s[sIdx] != t[tIdx] {
			return false
		}

		sIdx--
		tIdx--
	}

	return true
}

func backSpaceNextChar(str string, idx int) int {
	backSpaceCount := 0

	for idx >= 0 {
		if str[idx] == '#' {
			// add the backspace charactor count
			backSpaceCount++

			// skip the next charactor of '#'
			idx--
		} else if backSpaceCount > 0 {
			// subtract the backspace charactor count
			backSpaceCount--

			// skip the next charactor of '#'
			idx--
		} else {
			break
		}
	}

	return idx
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

func Test_backspaceCompare2(t *testing.T) {
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
			backspaceCompare2(tc.args.s, tc.args.t),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_backspaceCompare1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		backspaceCompare1("ab#c", "ad#c")
	}
}

func Benchmark_backspaceCompare2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		backspaceCompare2("ab#c", "ad#c")
	}
}