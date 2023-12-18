package medium

import (
	"fmt"
	"testing"

	"leetcode/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/container-with-most-water/

// method 1
// 1) use two pointers, leftIdx and rightIdx
// 2) use one for loop, to scan the height
// 3) calculate the currentArea, and compare with the maxArea
// 4) if height[leftIdx] < height[rightIdx], then leftIdx++
// 5) if height[leftIdx] >= height[rightIdx], then rightIdx--
// TC = O(N), SC = O(1)
func maxArea1(height []int) int {
	maxArea := 0
	leftIdx := 0
	rightIdx := len(height) - 1

	for leftIdx < rightIdx {
		currentArea := util.Min(height[leftIdx], height[rightIdx]) * (rightIdx - leftIdx)
		maxArea = util.Max(maxArea, currentArea)

		if height[leftIdx] < height[rightIdx] {
			leftIdx++
		} else {
			rightIdx--
		}
	}

	return maxArea
}

// method 2, this method without using the min() and max() function more faster
// logic thinking is same as method 1
// TC = O(N), SC = O(1)
// * this is the best solution for me currently
func maxArea2(height []int) int {
	maxArea := 0
	leftIdx := 0
	rightIdx := len(height) - 1

	for leftIdx < rightIdx {
		h := 0
		w := rightIdx - leftIdx

		if height[leftIdx] < height[rightIdx] {
			h = height[leftIdx]
			leftIdx++
		} else {
			h = height[rightIdx]
			rightIdx--
		}

		currentArea := h * w
		if currentArea > maxArea {
			maxArea = currentArea
		}
	}

	return maxArea
}

func Test_maxArea1(t *testing.T) {
	type args struct {
		height []int
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
				height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			},
			expected: expected{
				result: 49,
			},
		},
		{
			name: "2",
			args: args{
				height: []int{1, 1},
			},
			expected: expected{
				result: 1,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			maxArea1(tc.args.height),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_maxArea2(t *testing.T) {
	type args struct {
		height []int
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
				height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			},
			expected: expected{
				result: 49,
			},
		},
		{
			name: "2",
			args: args{
				height: []int{1, 1},
			},
			expected: expected{
				result: 1,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			maxArea2(tc.args.height),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
