package hard

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/basic-calculator/description/

// method 1 stack
// 1) use stack to store the result and curSign before '('
// 2) use curNum to store the current number
// 3) use curSign to store the current sign
// 4) use result to store the result of current calculation
// 5) when meet '(', store the result and curSign before '(' into stack
// 6) when meet ')', get the sign before '(' and the result before '(' from stack
// 7) when meet '+', '-', calculate the result and reset curSign and curNum
// 8) when meet number, calculate the curNum
// 9) when meet the end of string, calculate the result
// 10) return the result
// TC = O(N), SC = O(N)
func calculate1(s string) int {
	stack := []int{} // store the iterated result and curSign before '('
	curNum := 0      // current number
	curSign := 1     // curSign for current calculation
	result := 0

	for _, char := range s {
		// curNum get the current number of string, like: 123 + 1, curNum can get 123
		if char >= '0' && char <= '9' {
			curNum = curNum*10 + int(char-'0')
			continue
		}

		switch char {
		case '+':
			result += curSign * curNum

			// reset curSign and curNum for next calculation
			curSign = 1
			curNum = 0
		case '-':
			result += curSign * curNum

			// reset curSign and curNum for next calculation
			curSign = -1
			curNum = 0
		case '(':
			// store the result and curSign before '(' into stack
			// eg. 123 - (1 + 2), stack = [result, curSign] => [123, -1]
			stack = append(stack, result)
			stack = append(stack, curSign)

			// reset result and curSign for next calculation
			result = 0
			curSign = 1
		case ')':
			result += curSign * curNum

			// reset curSign and curNum for next calculation
			curSign = 1
			curNum = 0

			// get the sign before '('
			result *= stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop the sign

			// get the result before '('
			result += stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop the result
		}
	}

	// add the last number
	result += curSign * curNum

	return result
}

// method 2 recursion
// 1) use i to store the current index of string
// 2) use curNum to store the current number
// 3) use curSign to store the current sign
// 4) use result to store the result of current calculation
// 5) when meet '(', calculate the result in the parentheses
// 6) when meet ')', return the result
// 7) when meet '+', '-', calculate the result and reset curSign and curNum
// 8) when meet number, calculate the curNum
// 9) when meet the end of string, calculate the result
// 10) return the result
// TC = O(N), SC = O(N)
func calculate2(s string) int {
	i := 0

	return calculate2Helper(s, &i)
}

func calculate2Helper(s string, i *int) int {
	curNum := 0
	curSign := 1
	result := 0

	for *i < len(s) {
		char := s[*i]
		*i++

		if char >= '0' && char <= '9' {
			curNum = curNum*10 + int(char-'0')
			continue
		}

		switch char {
		case '+':
			result += curSign * curNum

			// reset curSign and curNum for next calculation
			curSign = 1
			curNum = 0
		case '-':
			result += curSign * curNum

			// reset curSign and curNum for next calculation
			curSign = -1
			curNum = 0
		case '(':
			// calculate the result in the parentheses
			result += curSign * calculate2Helper(s, i)

			// * no need to reset curSign and curNum for next calculation
			// curSign = 1
			// curNum = 0
		case ')':
			result += curSign * curNum

			// * no need to reset curSign and curNum for next calculation
			// curSign = 1
			// curNum = 0

			return result
		}
	}

	// add the last number
	result += curSign * curNum

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
				s: "1 + 1",
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "2",
			args: args{
				s: " 2-1 + 2 ",
			},
			expected: expected{
				result: 3,
			},
		},
		{
			name: "3",
			args: args{
				s: "(1+(4+5+2)-3)+(6+8)",
			},
			expected: expected{
				result: 23,
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

func Test_calculate2(t *testing.T) {
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
				s: "1 + 1",
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "2",
			args: args{
				s: " 2-1 + 2 ",
			},
			expected: expected{
				result: 3,
			},
		},
		{
			name: "3",
			args: args{
				s: "(1+(4+5+2)-3)+(6+8)",
			},
			expected: expected{
				result: 23,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			calculate2(tc.args.s),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_calculate1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculate1("1 + 1")
	}
}

func Benchmark_calculate2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculate2("1 + 1")
	}
}