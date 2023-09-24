package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/3sum/description/

// method 1
// use two for loop, to scan the nums
// 1st, saving the left multiply nums of current num[i] to the result[i]
// 2nd, saving the right multiply nums of current num[i] to the result[i]
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
func productExceptSelf1(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// 1st, calculate the left multiply nums of current num[i] and save the result to the result[i]
	leftMultiplyNum := 1
	for i := 0; i < n; i++ {
		result[i] = leftMultiplyNum
		leftMultiplyNum *= nums[i]
	}

	// 2nd, calculate the right multiply nums of current num[i] and the result will be multiply with the result[i] (this result[i] is the left multiply nums of current num[i]), then save the result to the result[i]
	rightMultiplyNum := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= rightMultiplyNum
		rightMultiplyNum *= nums[i]
	}

	return result
}

// method 2, this method is provide by my friend. This method with more logic thinking.
// use one for loop, to scan the nums, and use multiplyNums to store the multiply result of all nums
// 1st, if the current num[i] is 0, then we need to count the zeroCount, and skip the current loop
// 2nd, if the current num[i] is not 0, then we need to divide the multiplyNums with current num[i], and save the result to the result[i]
// TC = O(N), SC = O(N)
func productExceptSelf2(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	multiplyNums := 1
	zeroCount := 0

	for i := 0; i < n; i++ {
		// this step is essential, if the current num[i] is 0, then we need to count the zeroCount, and skip the current loop, because this will make the multiplyNums = 0
		if nums[i] == 0 {
			zeroCount++
			continue
		}
		multiplyNums *= nums[i]
	}

	for i := 0; i < n; i++ {
		if nums[i] == 0 && zeroCount == 1 {
			result[i] = multiplyNums
		} else if zeroCount >= 1 {
			result[i] = 0
		} else {
			// this step is essential, if the current num[i] is not 0, then we need to divide the multiplyNums with current num[i], and save the result to the result[i]
			result[i] = multiplyNums / nums[i]
		}
	}

	return result
}

func Test_productExceptSelf1(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{1, 2, 3, 4},
			},
			expected: expected{
				result: []int{24, 12, 8, 6},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{-1, 1, 0, -3, 3},
			},
			expected: expected{
				result: []int{0, 0, 9, 0, 0},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			productExceptSelf1(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_productExceptSelf2(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{1, 2, 3, 4},
			},
			expected: expected{
				result: []int{24, 12, 8, 6},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{-1, 1, 0, -3, 3},
			},
			expected: expected{
				result: []int{0, 0, 9, 0, 0},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			productExceptSelf2(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
