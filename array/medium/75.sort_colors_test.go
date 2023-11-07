package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sort-colors/description/

// method 1 Dutch National Flag Algorithm
// 1) use three pointers, leftIndex, rightIndex, currentIndex
// 2) use one for loop, to scan the nums
// 3) if nums[currentIndex] == 0, then swap the nums[currentIndex] with nums[leftIndex], and leftIndex++ and currentIndex++
// 4) if nums[currentIndex] == 2, then swap the nums[currentIndex] with nums[rightIndex], and rightIndex--
// 5) if nums[currentIndex] == 1, then currentIndex++
// TC = O(N), SC = O(1)
// * this is the best solution for me currently
func sortColors1(nums []int) {
	leftIndex := 0
	rightIndex := len(nums) - 1
	currentIndex := 0

	for currentIndex <= rightIndex {
		if nums[currentIndex] == 0 {
			// swap the 0 to the left
			nums[currentIndex], nums[leftIndex] = nums[leftIndex], nums[currentIndex]
			leftIndex++
			currentIndex++
		} else if nums[currentIndex] == 2 {
			// swap the 2 to the right
			nums[currentIndex], nums[rightIndex] = nums[rightIndex], nums[currentIndex]
			rightIndex--
		} else {
			currentIndex++
		}
	}

	// this method have repeat step i-- & i++
	/*
		for i:=0; i<=rightIndex; i++ {
			if nums[i] == 2 {
				nums[i], nums[rightIndex] = nums[rightIndex], nums[i]
				i--
				rightIndex--
			} else if nums[i] == 0 {
				nums[i], nums[leftIndex] = nums[leftIndex], nums[i]
				leftIndex++
			}
		}
	*/
}

func Test_sortColors1(t *testing.T) {
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
				nums: []int{2, 0, 2, 1, 1, 0},
			},
			expected: expected{
				result: []int{0, 0, 1, 1, 2, 2},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{2, 0, 1},
			},
			expected: expected{
				result: []int{0, 1, 2},
			},
		},
	}

	for _, tc := range testCases {
		sortColors1(tc.args.nums)
		assert.Equal(
			t,
			tc.expected.result,
			tc.args.nums,
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
