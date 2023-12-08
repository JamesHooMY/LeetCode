package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/daily-temperatures/description/

// method 1 stack monotonic decreasing
// 1) use a stack to store the index of iterated temperature in temperatures slice
// 2) if the current temperature is greater than the top of stack, then pop the top of stack
// 3) the waiting days of the top of stack is the difference between the current index and the top of stack
// 4) push the current index into stack
// 5) finally, the waiting days of remaining index in stack is 0
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
func dailyTemperatures1(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)

	// monotonic decreasing stack
	stack := []int{} // value store the index of iterated temperature in temperatures slice

	for i, temp := range temperatures {
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]

			// save the waiting days of the top of stack
			result[top] = i - top

			// pop the top of stack, the temperature of top is less than the current temperature
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return result
}

func Test_dailyTemperatures1(t *testing.T) {
	type args struct {
		temperatures []int
	}
	type expected struct {
		result []int
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
				temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			},
			expected: expected{
				result: []int{1, 1, 4, 2, 1, 1, 0, 0},
			},
		},
		{
			name: "2",
			args: args{
				temperatures: []int{30, 40, 50, 60},
			},
			expected: expected{
				result: []int{1, 1, 1, 0},
			},
		},
		{
			name: "3",
			args: args{
				temperatures: []int{30, 60, 90},
			},
			expected: expected{
				result: []int{1, 1, 0},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			dailyTemperatures1(tc.args.temperatures),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
