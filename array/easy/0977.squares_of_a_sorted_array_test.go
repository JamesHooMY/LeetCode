package easy

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/squares-of-a-sorted-array/description/

// method 1 brute force
// 1) use two for loop, first for loop square the num of nums
// 2) use sort package to sort the new nums array
// TC = O(NlogN), SC = O(logN)
func sortedSquares1(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}

	// use sort package to sort the new nums array
	sort.Ints(nums)

	return nums
}

// method 2 two pointer
// 1) use two pointer, one pointer is leftIndex, another pointer is rightIndex; one sortedSquaresList to store the result
// 2) use one for loop, the index start form the last of sortedSquaresList, if leftValue^2 > rightValue^2, then add leftValue^2 to sortedSquaresList[i], and leftIndex++
// 3) else add rightValue^2 to sortedSquaresList[i], and rightIndex--
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
func sortedSquares2(nums []int) []int{
	leftIndex := 0
    rightIndex := len(nums) - 1

    sortedSquaresList := make([]int, len(nums))

    for i := len(nums) - 1; i >= 0; i-- {
        // these two variable must initialize in for loop
        leftValue := nums[leftIndex]
        rightValue := nums[rightIndex]

        if leftValue * leftValue > rightValue * rightValue {
            sortedSquaresList[i] = leftValue * leftValue
            leftIndex++
        } else {
            sortedSquaresList[i] = rightValue * rightValue
            rightIndex--
        }
    }

    return sortedSquaresList
}


func Test_sortedSquares1(t *testing.T) {
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
				nums: []int{-4, -1, 0, 3, 10},
			},
			expected: expected{
				result: []int{0, 1, 9, 16, 100},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{-7, -3, 2, 3, 11},
			},
			expected: expected{
				result: []int{4, 9, 9, 49, 121},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			sortedSquares1(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_sortedSquares2(t *testing.T) {
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
				nums: []int{-4, -1, 0, 3, 10},
			},
			expected: expected{
				result: []int{0, 1, 9, 16, 100},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{-7, -3, 2, 3, 11},
			},
			expected: expected{
				result: []int{4, 9, 9, 49, 121},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			sortedSquares2(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}