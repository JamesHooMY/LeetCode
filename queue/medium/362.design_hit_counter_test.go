package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://github.com/doocs/leetcode/blob/main/solution/0300-0399/0362.Design%20Hit%20Counter/README_EN.md

type HitCounter struct {
	queueTimestamp []int
}

func ConstructorHitCounter() HitCounter {
	return HitCounter{
		queueTimestamp: []int{},
	}
}

func (this *HitCounter) Hit(timestamp int) {
	this.queueTimestamp = append(this.queueTimestamp, timestamp)
}

func (this *HitCounter) GetHits(timestamp int) int {
	// Remove all timestamp that is 5 seconds ago, [1, 2, 3, 4, 5] -> Hit 6 -> [2, 3, 4, 5, 6]
	for len(this.queueTimestamp) > 0 && timestamp-this.queueTimestamp[0] >= 5 {
		this.queueTimestamp = this.queueTimestamp[1:]
	}

	return len(this.queueTimestamp)
}

func Test_HitCounter_GetHits(t *testing.T) {
	hitCounter := ConstructorHitCounter()

	type args struct {
		timestamp int
	}
	type expected struct {
		result int
	}
	type testCase struct {
		name           string
		args           args
		hitCounterFunc func(hitCounter *HitCounter, timestamp int)
		expected       expected
	}

	testCases := []testCase{
		{
			name: "1",
			args: args{
				timestamp: 1,
			},
			hitCounterFunc: func(hitCounter *HitCounter, timestamp int) {
				hitCounter.Hit(timestamp)
			},
			expected: expected{
				result: 1,
			},
		},
		{
			name: "2",
			args: args{
				timestamp: 2,
			},
			hitCounterFunc: func(hitCounter *HitCounter, timestamp int) {
				hitCounter.Hit(timestamp)
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "3",
			args: args{
				timestamp: 3,
			},
			hitCounterFunc: func(hitCounter *HitCounter, timestamp int) {
				hitCounter.Hit(timestamp)
			},
			expected: expected{
				result: 3,
			},
		},
		{
			name: "4",
			args: args{
				timestamp: 4,
			},
			hitCounterFunc: func(hitCounter *HitCounter, timestamp int) {
				hitCounter.Hit(timestamp)
			},
			expected: expected{
				result: 4,
			},
		},
		{
			name: "5",
			args: args{
				timestamp: 5,
			},
			hitCounterFunc: func(hitCounter *HitCounter, timestamp int) {
				hitCounter.Hit(timestamp)
			},
			expected: expected{
				result: 5,
			},
		},
		{
			name: "6",
			args: args{
				timestamp: 6,
			},
			hitCounterFunc: func(hitCounter *HitCounter, timestamp int) {
				hitCounter.Hit(timestamp)
			},
			expected: expected{
				result: 5,
			},
		},
		{
			name: "7",
			args: args{
				timestamp: 7,
			},
			hitCounterFunc: func(hitCounter *HitCounter, timestamp int) {
				hitCounter.Hit(timestamp)
			},
			expected: expected{
				result: 5,
			},
		},
	}

	for _, tc := range testCases {
		tc.hitCounterFunc(&hitCounter, tc.args.timestamp)

		assert.Equal(
			t,
			tc.expected.result,
			hitCounter.GetHits(tc.args.timestamp),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
