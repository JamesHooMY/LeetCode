package hard

import (
	"fmt"
	"testing"

	"leetcode/array/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/longest-valid-parentheses/description/

// method 1 stack
// 1) use stackIdx to store the index of '('
// 2) use one for loop, to scan the string
// 3) if char == '(', push the index into stackIdx
// 4) if char == ')', pop the index from stackIdx
// 5) if stackIdx is empty, push the index into stackIdx
// 6) if stackIdx is not empty, calculate the length of valid parentheses
// 7) finally, return the max length of valid parentheses
// TC = O(N), SC = O(N)
func longestValidParentheses1(s string) int {
	if len(s) < 2 {
		return 0
	}

	stackIdx := []int{-1} // store the left edge index of valid parentheses
	maxLength := 0

	for i, char := range s {
		if char == '(' {
			stackIdx = append(stackIdx, i)
		} else {
			// default stackIdx = [-1], this pop will not cause error if the string start with ')'
			stackIdx = stackIdx[:len(stackIdx)-1]

			if len(stackIdx) == 0 {
				// store the index of current ')' as the left edge of valid parentheses at the bottom of stackIdx for next calculation
				stackIdx = append(stackIdx, i)
				continue
			}

			curLength := i - stackIdx[len(stackIdx)-1]
			maxLength = util.Max(maxLength, curLength)
		}
	}

	return maxLength
}

// method 2 dynamic programming with one slice
// 1) use sliceLength to store the length of valid parentheses which end with s[i] which is ')'
// 2) use one for loop, to scan the string
// 3) if char == '(', sliceLength[i] = 0
// 4) if char == ')', if s[i-1] == '(', sliceLength[i] = 2 + sliceLength[i-2]
// 5) if char == ')', if s[i-1] == ')' && s[i-sliceLength[i-1]-1] == '(', sliceLength[i] = sliceLength[i-1] + 2 + sliceLength[i-sliceLength[i-1]-2]
// 6) finally, return the max length of valid parentheses
// TC = O(N), SC = O(N)
func longestValidParentheses2(s string) int {
	if len(s) < 2 {
		return 0
	}

	sliceLength := make([]int, len(s)) // sliceLength[i] is the length of valid parentheses which end with s[i] which is ')'
	maxLength := 0

	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				// "()" length = 2
				sliceLength[i] = 2

				// "()()" length = 2 + sliceLength[i-2]
				if i-2 >= 0 {
					sliceLength[i] += sliceLength[i-2]
				}
			} else if i-sliceLength[i-1]-1 >= 0 && s[i-sliceLength[i-1]-1] == '(' {
				// "((...))" length = sliceLength[i-1] + 2
				sliceLength[i] = sliceLength[i-1] + 2

				// "()((...))" length = sliceLength[i-1] + 2 + sliceLength[i-sliceLength[i-1]-2]
				if i-sliceLength[i-1]-2 >= 0 {
					sliceLength[i] += sliceLength[i-sliceLength[i-1]-2]
				}
			}
		}

		maxLength = util.Max(maxLength, sliceLength[i])
	}

	return maxLength
}

// method 3 two counters, this method is slower than method 2 if the string is extremely long !!! due to the two for loop, TC = O(2N)
// 1) use leftCounter and rightCounter to store the number of '(' and ')'
// 2) use two for loop, to scan the string from left to right and from right to left
// 3) if char == '(', leftCounter++, else rightCounter++
// 4) if leftCounter == rightCounter, calculate the length of valid parentheses
// 5) if leftCounter < rightCounter, reset leftCounter and rightCounter to 0
// 6) finally, return the max length of valid parentheses
// TC = O(N), SC = O(1)
// * this is the best solution for me currently
func longestValidParentheses3(s string) int {
	if len(s) < 2 {
		return 0
	}

	leftCounter, rightCounter := 0, 0 // leftCounter is the number of '(', rightCounter is the number of ')'
	maxLength := 0

	// scan from leftCounter to rightCounter
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			leftCounter++
		} else {
			rightCounter++
		}

		if leftCounter == rightCounter {
			maxLength = util.Max(maxLength, leftCounter*2)
		} else if leftCounter < rightCounter {
			// ')' is more than '(', reset leftCounter and rightCounter to 0
			leftCounter, rightCounter = 0, 0
		}
	}

	// scan from rightCounter to leftCounter, need one more for loop to handle the case: s = "(()"
	leftCounter, rightCounter = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			rightCounter++
		} else {
			leftCounter++
		}

		if rightCounter == leftCounter {
			maxLength = util.Max(maxLength, rightCounter*2)
		} else if rightCounter < leftCounter {
			leftCounter, rightCounter = 0, 0
		}
	}

	return maxLength
}

func Test_longestValidParentheses1(t *testing.T) {
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
				s: "(()",
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "2",
			args: args{
				s: ")()())",
			},
			expected: expected{
				result: 4,
			},
		},
		{
			name: "3",
			args: args{
				s: "",
			},
			expected: expected{
				result: 0,
			},
		},
		{
			name: "4",
			args: args{
				s: "()()()",
			},
			expected: expected{
				result: 6,
			},
		},
		{
			name: "5",
			args: args{
				s: "(()()",
			},
			expected: expected{
				result: 4,
			},
		},
		{
			name: "6",
			args: args{
				s: "()(())",
			},
			expected: expected{
				result: 6,
			},
		},
		{
			name: "7",
			args: args{
				s: "(()())",
			},
			expected: expected{
				result: 6,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			longestValidParentheses1(tc.args.s),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_longestValidParentheses2(t *testing.T) {
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
				s: "(()",
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "2",
			args: args{
				s: ")()())",
			},
			expected: expected{
				result: 4,
			},
		},
		{
			name: "3",
			args: args{
				s: "",
			},
			expected: expected{
				result: 0,
			},
		},
		{
			name: "4",
			args: args{
				s: "()()()",
			},
			expected: expected{
				result: 6,
			},
		},
		{
			name: "5",
			args: args{
				s: "(()()",
			},
			expected: expected{
				result: 4,
			},
		},
		{
			name: "6",
			args: args{
				s: "()(())",
			},
			expected: expected{
				result: 6,
			},
		},
		{
			name: "7",
			args: args{
				s: "(()())",
			},
			expected: expected{
				result: 6,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			longestValidParentheses2(tc.args.s),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_longestValidParentheses3(t *testing.T) {
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
				s: "(()",
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "2",
			args: args{
				s: ")()())",
			},
			expected: expected{
				result: 4,
			},
		},
		{
			name: "3",
			args: args{
				s: "",
			},
			expected: expected{
				result: 0,
			},
		},
		{
			name: "4",
			args: args{
				s: "()()()",
			},
			expected: expected{
				result: 6,
			},
		},
		{
			name: "5",
			args: args{
				s: "(()()",
			},
			expected: expected{
				result: 4,
			},
		},
		{
			name: "6",
			args: args{
				s: "()(())",
			},
			expected: expected{
				result: 6,
			},
		},
		{
			name: "7",
			args: args{
				s: "(()())",
			},
			expected: expected{
				result: 6,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			longestValidParentheses3(tc.args.s),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_longestValidParentheses1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestValidParentheses1("(()")
	}
}

func Benchmark_longestValidParentheses2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestValidParentheses2("(()")
	}
}

func Benchmark_longestValidParentheses3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestValidParentheses3("(()")
	}
}
