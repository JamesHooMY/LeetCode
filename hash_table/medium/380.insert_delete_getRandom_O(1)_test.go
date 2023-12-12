package medium

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/insert-delete-getrandom-o1/description/

// method 1
type RandomizedSet struct {
	nums   []int
	numIdxMap map[int]int // key: num, value: index of num in nums
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums:   []int{},
		numIdxMap: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exist := this.numIdxMap[val]; exist {
		return false
	}

	this.nums = append(this.nums, val)
	this.numIdxMap[val] = len(this.nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, exist := this.numIdxMap[val]
	if !exist {
		return false
	}

	lastIndex := len(this.nums) - 1
	lastNum := this.nums[lastIndex]

	// change the lastNum index to the index of val, update the numIdxMap of lastNum
	this.nums[index] = lastNum
	this.numIdxMap[lastNum] = index

	// remove lastNum index in nums and delete val in numIdxMap
	this.nums = this.nums[:lastIndex]
	delete(this.numIdxMap, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	// get the value with a random index
	return this.nums[rand.Intn(len(this.nums))]
}

func Test_RandomizedSet_GetRandom(t *testing.T) {
	randomizedSet := Constructor()

	type args struct {
		val int
	}
	type expected struct {
		result int
	}
	type testCase struct {
		name                 string
		args                 args
		randomizedSetMethods func(randomizedSet *RandomizedSet, val int)
		expected             expected
	}

	testCases := []testCase{
		{
			name: "1",
			args: args{
				val: 1,
			},
			randomizedSetMethods: func(randomizedSet *RandomizedSet, val int) {
				randomizedSet.Insert(val)
			},
			expected: expected{
				result: 1,
			},
		},
		{
			name: "2",
			args: args{
				val: 2,
			},
			randomizedSetMethods: func(randomizedSet *RandomizedSet, val int) {
				randomizedSet.Remove(1)
				randomizedSet.Insert(val)
			},
			expected: expected{
				result: 2,
			},
		},
	}

	for _, tc := range testCases {
		tc.randomizedSetMethods(&randomizedSet, tc.args.val)

		assert.Equal(
			t,
			tc.expected.result,
			randomizedSet.GetRandom(),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
