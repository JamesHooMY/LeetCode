package hard

import (
	"fmt"
	"testing"

	"leetcode/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/largest-rectangle-in-histogram/description/

// method 1 two slices dynamic programming
// 1) use two slices, left and right, to store the first index left and right that has height lower than heights[i]
// 2) use one for loop, to scan the height from left to right, and alway keep the first index left that has height lower than heights[i]
// 3) use one for loop, to scan the height from right to left, and alway keep the first index right that has height lower than heights[i]
// 4) use one for loop, to scan the height, and calculate the result
// TC = O(N), SC = O(N)
func largestRectangleArea1(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	/*
		| index   | 0 | 1 | 2 | 3 | 4 | 5 |
		| heights | 2 | 1 | 5 | 6 | 2 | 3 |
		|---------|---|---|---|---|---|---|
		| left    | -1| -1|  1|  2|  1|  4|
		| right   |  1|  6|  4|  4|  6|  6|
	*/

	// store the first index left that has height lower than heights[i] for getting left farthest extend edge
	left := make([]int, n)

	// store the first index right that has height lower than heights[i] for getting right farthest extend edge
	right := make([]int, n)

	left[0] = -1 // the left edge of index 0 is -1，calculate in right[i] - left[i] - 1
	for i := 1; i < n; i++ {
		prevIndex := i - 1
		for prevIndex >= 0 && heights[prevIndex] >= heights[i] {
			// heights[prevIndex] >= heights[i], mean the left edge of prevIndex must be the left edge of i
			prevIndex = left[prevIndex]
		}
		left[i] = prevIndex
	}

	right[n-1] = n // the right edge of index n-1 is n，calculate in right[i] - left[i] - 1
	for i := n - 2; i >= 0; i-- {
		prevIndex := i + 1
		for prevIndex < n && heights[prevIndex] >= heights[i] {
			// heights[prevIndex] >= heights[i], mean the right edge of prevIndex must be the right edge of i
			prevIndex = right[prevIndex]
		}
		right[i] = prevIndex
	}

	maxArea := 0
	for i := 0; i < n; i++ {
		h := heights[i]
		w := right[i] - left[i] - 1

		maxArea = util.Max(maxArea, h*w)
	}

	// * this optimize method save the space right slice, but it cannot pass leetcode ! because of time limit exceeded
	// maxArea := 0
	// for i := 0; i < n; i++ {
	// 	// find the right farthest extend edge
	// 	rightIdx := i + 1
	// 	for rightIdx < n && heights[rightIdx] >= heights[i] {
	// 		rightIdx++
	// 	}

	// 	h := heights[i]
	// 	w := rightIdx - left[i] - 1
	// 	maxArea = util.Max(maxArea, h*w)
	// }

	return maxArea
}

// method 2 two pointers dynamic programming, cannot pass leetcode ! because of time limit exceeded
// 1) use two pointers, leftIdx and rightIdx
// 2) use one for loop, to scan the height from left to right, and alway keep the first index left that has height lower than heights[i]
// 3) use one for loop, to scan the height from right to left, and alway keep the first index right that has height lower than heights[i]
// 4) use one for loop, to scan the height, and calculate the result
// TC = O(N), SC = O(1)
func largestRectangleArea2(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	/*
		| index   | 0 | 1 | 2 | 3 | 4 | 5 |
		| heights | 2 | 1 | 5 | 6 | 2 | 3 |
		|---------|---|---|---|---|---|---|
		| left    | -1| -1|  1|  2|  1|  4|
		| right   |  1|  6|  4|  4|  6|  6|
	*/
	maxArea := 0
	for i := 0; i < n; i++ {
		// find the left farthest extend edge
		leftIdx := i - 1
		for leftIdx >= 0 && heights[leftIdx] >= heights[i] {
			leftIdx--
		}

		// find the right farthest extend edge
		rightIdx := i + 1
		for rightIdx < n && heights[rightIdx] >= heights[i] {
			rightIdx++
		}

		h := heights[i]
		w := rightIdx - leftIdx - 1
		maxArea = util.Max(maxArea, h*w)
	}

	return maxArea
}

// method 3 stackIdx monotonous increasing
// 1) use a stackIdx to store the index of iterated height in heights slice
// 2) if the current height is greater than the top of stackIdx, then pop the top of stackIdx
// 3) the area of the top of stackIdx is the height of top * (the difference between the current index and the top of stackIdx)
// 4) push the current index into stackIdx
// 5) finally, the area of remaining index in stackIdx is the height of top * (the difference between the current index and the top of stackIdx)
// TC = O(N), SC = O(N)
func largestRectangleArea3(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	maxArea := 0
	stackIdx := []int{} // value store the index of iterated height in heights slice

	for i := 0; i <= n; i++ {
		curHeight := 0
		if i < n {
			curHeight = heights[i]
		}

		/*
			[2,5,6,1]

			i = 0, stackIdx = [], maxArea = 0, curHeight = 2 => stackIdx = [0], maxArea = 0
			i = 1, stackIdx = [0], maxArea = 0, curHeight = 5 => stackIdx = [0, 1], maxArea = 0
			i = 2, stackIdx = [0, 1], maxArea = 0, curHeight = 6 => stackIdx = [0, 1, 2], maxArea = 0

			i = 3, stackIdx = [0, 1, 2], maxArea = 0, curHeight = 1
			iterating stackIdx:
			=> stackIdx = [0, 1], width = 1, maxArea = 6
			=> stackIdx = [0], width = 2, maxArea = 10
			=> stackIdx = [], width = 3, maxArea = 10
		*/
		for len(stackIdx) > 0 && curHeight < heights[stackIdx[len(stackIdx)-1]] {
			topIndex := stackIdx[len(stackIdx)-1]
			stackIdx = stackIdx[:len(stackIdx)-1] // pop the top index

			// calculate the area of water of topIndex position
			// if stackIdx is empty, then the width of topIndex position is i
			// if stackIdx is not empty, then the width of topIndex position is i - stackIdx[len(stackIdx)-1] - 1
			width := i
			if len(stackIdx) > 0 {
				width = i - stackIdx[len(stackIdx)-1] - 1
			}

			maxArea = util.Max(maxArea, heights[topIndex]*width)
		}

		stackIdx = append(stackIdx, i)
	}

	return maxArea
}

func Test_largestRectangleArea1(t *testing.T) {
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
				height: []int{2, 1, 5, 6, 2, 3},
			},
			expected: expected{
				result: 10,
			},
		},
		{
			name: "2",
			args: args{
				height: []int{2, 1, 2},
			},
			expected: expected{
				result: 3,
			},
		},
		{
			name: "3",
			args: args{
				height: []int{1, 2, 2},
			},
			expected: expected{
				result: 4,
			},
		},
		{
			name: "4",
			args: args{
				height: []int{999, 999, 999, 999},
			},
			expected: expected{
				result: 3996,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			largestRectangleArea1(tc.args.height),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_largestRectangleArea2(t *testing.T) {
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
				height: []int{2, 1, 5, 6, 2, 3},
			},
			expected: expected{
				result: 10,
			},
		},
		{
			name: "2",
			args: args{
				height: []int{2, 1, 2},
			},
			expected: expected{
				result: 3,
			},
		},
		{
			name: "3",
			args: args{
				height: []int{1, 2, 2},
			},
			expected: expected{
				result: 4,
			},
		},
		{
			name: "4",
			args: args{
				height: []int{999, 999, 999, 999},
			},
			expected: expected{
				result: 3996,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			largestRectangleArea2(tc.args.height),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_largestRectangleArea3(t *testing.T) {
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
				height: []int{2, 1, 5, 6, 2, 3},
			},
			expected: expected{
				result: 10,
			},
		},
		{
			name: "2",
			args: args{
				height: []int{2, 1, 2},
			},
			expected: expected{
				result: 3,
			},
		},
		{
			name: "3",
			args: args{
				height: []int{1, 2, 2},
			},
			expected: expected{
				result: 4,
			},
		},
		{
			name: "4",
			args: args{
				height: []int{999, 999, 999, 999},
			},
			expected: expected{
				result: 3996,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			largestRectangleArea3(tc.args.height),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

var input = make([]int, 100000)

// benchmark
func Benchmark_largestRectangleArea1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		largestRectangleArea1(input)
	}
}

func Benchmark_largestRectangleArea2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		largestRectangleArea2(input)
	}
}

func Benchmark_largestRectangleArea3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		largestRectangleArea3(input)
	}
}
