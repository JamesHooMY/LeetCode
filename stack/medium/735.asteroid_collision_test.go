package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/asteroid-collision/description/

// method 1 stack
// 1) use a stack to store the result of asteroid collision
// 2) if the top of stack is positive and the current asteroid is negative, then asteroid collision
// 3) if the top of stack is negative and the current asteroid is positive, then no asteroid collision
// 4) if the top of stack is positive and the current asteroid is positive, then no asteroid collision
// 5) if the top of stack is negative and the current asteroid is negative, then no asteroid collision
// 6) finally, the stack is the result of asteroid collision
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
func asteroidCollision1(asteroids []int) []int {
	// store the result of asteroid collision
	stack := []int{}

	for _, asteroid := range asteroids {
		// * asteroid collision, when the top of stack is positive and the current asteroid is negative
		for len(stack) > 0 && stack[len(stack)-1] > 0 && asteroid < 0 {
			if stack[len(stack)-1] < -asteroid {
				// * the top of stack is destroyed
				stack = stack[:len(stack)-1]
				continue
			} else if stack[len(stack)-1] == -asteroid {
				// * the top of stack and the current asteroid are destroyed
				stack = stack[:len(stack)-1]
				// * the current asteroid is destroyed
				asteroid = 0
			} else {
				// condition stack[len(stack)-1] > -asteroid
				// * the current asteroid is destroyed
				asteroid = 0
			}

			// after condition stack[len(stack)-1] == -asteroid or condition stack[len(stack)-1] > -asteroid
			break
		}

		// this two asteroid [8,-8], [-2,-1,1,2], will get empty stack in the end, so we can not use this condition !!!
        // if !(len(stack) > 0 && stack[len(stack)-1] > 0 && asteroid < 0) {
        //     stack = append(stack, asteroid)
        // }

		if asteroid != 0 {
			stack = append(stack, asteroid)
		}
	}

	return stack
}

func Test_asteroidCollision1(t *testing.T) {
	type args struct {
		asteroids []int
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
				asteroids: []int{5, 10, -5},
			},
			expected: expected{
				result: []int{5, 10},
			},
		},
		{
			name: "2",
			args: args{
				asteroids: []int{8, -8},
			},
			expected: expected{
				result: []int{},
			},
		},
		{
			name: "3",
			args: args{
				asteroids: []int{10, 2, -5},
			},
			expected: expected{
				result: []int{10},
			},
		},
		{
			name: "4",
			args: args{
				asteroids: []int{-2, -1, 1, 2},
			},
			expected: expected{
				result: []int{-2, -1, 1, 2},
			},
		},
		{
			name: "5",
			args: args{
				asteroids: []int{-2, -2, 1, -2},
			},
			expected: expected{
				result: []int{-2, -2, -2},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			asteroidCollision1(tc.args.asteroids),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
