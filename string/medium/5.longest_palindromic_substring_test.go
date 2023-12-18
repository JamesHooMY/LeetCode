package medium

import (
	"fmt"
	"math"
	"testing"

	"leetcode/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/longest-palindromic-substring/description/

// method 1 center expand, 2 pointers
// 1) for each char in s, expand around center
// 2) return the longest palindrome
// TC = O(N^2), SC = O(1)
// * this is the best solution for me currently
func longestPalindrome1(s string) string {
	if len(s) < 2 {
		return s
	}

	var (
		start int
		// end    int
		maxLength int
	)

	for i := 0; i < len(s); i++ {
		oddLen := expandAroundCenter(s, i, i)
		evenLen := expandAroundCenter(s, i, i+1)
		maxLen := util.Max(oddLen, evenLen)

		// if maxLen > end-start {
		// 	start = i - (maxLen-1)/2
		// 	end = i + maxLen/2
		// }
		if maxLen > maxLength {
			start = i - (maxLen-1)/2
			maxLength = maxLen
		}
	}

	// return s[start : end+1]
	return s[start : start+maxLength]
}

func expandAroundCenter(s string, left int, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return right - left - 1
}

// method 2  Manacher's Algorithm, hard to understand for me currently, the performance not better than method 1 (slower than method 1, even though the TC is O(N), the s length was tested with math.MaxInt32 / 100000)
// 1) insert a special char between each char in s, e.g. "babad" -> "#b#a#b#a#d#" (process s to sb)
// 2) expand around center for each char in sb
// 3) return the longest palindrome
// TC = O(N), SC = O(N)
func longestPalindrome2(s string) string {
	if len(s) < 2 {
		return s
	}

	// insert a special char between each char in s, this process is to make the length of sb is odd
	// e.g. "babad" -> "#b#a#b#a#d#"
	var sb []byte
	sb = append(sb, '#')
	for i := 0; i < len(s); i++ {
		sb = append(sb, s[i])
		sb = append(sb, '#')
	}

	// p[i] is the 1/2 length of the longest palindrome with center i
	p := make([]int, len(sb))
	var (
		center   int
		maxRight int
		maxLen   int
		begin    int
	)

	/*
		sb: # b # a # b # a # d #
		i:  0 1 2 3 4 5 6 7 8 9 10
		p:  0 0 0 0 0 0 0 0 0 0 0
	*/
	for i := 0; i < len(sb); i++ {
		/*
			i:=0, maxRight:=0, skip
			i:=1, maxRight:=0, skip
			i:=2, maxRight:=2, skip
			i:=3, maxRight:=2, skip
			i:=4, maxRight:=6, center:=3, p[4]:=0, p[2]:=0 => mirror:=2*3-4=2, p[4]:=min(6-4, p[2])=0
			i:=5, maxRight:=6, center:=3, p[5]:=0, p[1]:=0 => mirror:=2*3-5=1, p[5]:=min(6-5, p[1])=0
			i:=6, maxRight:=8, center:=5, p[6]:=0, p[4]:=0 => mirror:=2*5-6=4, p[6]:=min(8-6, p[4])=0
			i:=7, maxRight:=8, center:=5, p[7]:=0, p[3]:=3 => mirror:=2*5-7=3, p[7]:=min(8-7, p[3])=1
		*/
		if i < maxRight {
			mirror := 2*center - i
			p[i] = util.Min(maxRight-i, p[mirror])
		}

		/*
			i:=0, left:=-1, right:=1
			i:=1, left:=0, right:=2
			i:=2, left:=1, right:=3
			i:=3, left:=2, right:=4
			i:=4, left:=3, right:=5
			i:=5, left:=4, right:=6
			i:=6, left:=5, right:=7
			i:=7, left:=5, right:=9
		*/
		left := i - (p[i] + 1)
		right := i + (p[i] + 1)

		/*
			i:=0, left:=-1, skip
			i:=1, left:=0, right:=2, sb[left] == sb[right] => p[1]:=1, left:=-1, right:=3
			i:=2, left:=1, right:=3, sb[left] != sb[right], skip
			i:=3, left:=2, right:=4, sb[left] == sb[right] => => => p[3]:=3, left:=-1, right:=7
			i:=4, left:=3, right:=5, sb[left] != sb[right], skip
			i:=5, left:=4, right:=6, sb[left] == sb[right] => => => p[5]:=3, left:=1, right:=9
			i:=6, left:=5, right:=7, sb[left] != sb[right], skip
			i:=7, left:=5, right:=9, sb[left] != sb[right], skip
		*/
		for left >= 0 && right < len(sb) && sb[left] == sb[right] {
			p[i]++
			left--
			right++
		}

		/*
			i:=0, i+p[0]:=0, maxRight:=0, skip
			i:=1, i+p[1]:=2, maxRight:=0 => center:=1, maxRight:=2
			i:=2, i+p[2]:=2, maxRight:=2, skip
			i:=3, i+p[3]:=6, maxRight:=2 => center:=3, maxRight:=6
			i:=4, i+p[4]:=4, maxRight:=6, skip
			i:=5, i+p[5]:=8, maxRight:=6 => center:=5, maxRight:=8
			i:=6, i+p[6]:=6, maxRight:=8, skip
			i:=7, i+p[7]:=8, maxRight:=8, skip
		*/
		if i+p[i] > maxRight {
			center = i
			maxRight = i + p[i]
		}

		/*
			i:=0, p[0]:=0, maxLen:=0, skip
			i:=1, p[1]:=1, maxLen:=0 => maxLen:=1, begin:=(1-1)/2=0
			i:=2, p[2]:=0, maxLen:=1, skip
			i:=3, p[3]:=3, maxLen:=1 => maxLen:=3, begin:=(3-3)/2=0
			i:=4, p[4]:=0, maxLen:=3, skip
			i:=5, p[5]:=3, maxLen:=3, skip
			i:=6, p[6]:=0, maxLen:=3, skip
			i:=7, p[7]:=1, maxLen:=3, skip
		*/
		if p[i] > maxLen {
			maxLen = p[i]
			begin = (i - maxLen) / 2
		}
	}

	return s[begin : begin+maxLen]
}

func Test_longestPalindrome1(t *testing.T) {
	type args struct {
		s string
	}
	type expected struct {
		result string
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
				s: "babad",
			},
			expected: expected{
				result: "bab",
			},
		},
		{
			name: "2",
			args: args{
				s: "cbbd",
			},
			expected: expected{
				result: "bb",
			},
		},
		{
			name: "3",
			args: args{
				s: "a",
			},
			expected: expected{
				result: "a",
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			longestPalindrome1(tc.args.s),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_longestPalindrome2(t *testing.T) {
	type args struct {
		s string
	}
	type expected struct {
		result string
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
				s: "babad",
			},
			expected: expected{
				result: "bab",
			},
		},
		{
			name: "2",
			args: args{
				s: "cbbd",
			},
			expected: expected{
				result: "bb",
			},
		},
		{
			name: "3",
			args: args{
				s: "a",
			},
			expected: expected{
				result: "a",
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			longestPalindrome2(tc.args.s),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func result() string {
	result := ""
	len := math.MaxInt32 / 100000

	for i := 0; i < len; i++ {
		result += "babad"
	}
	return result
}

// benchmark
func Benchmark_longestPalindrome1(b *testing.B) {
	r := result()
	for i := 0; i < b.N; i++ {
		longestPalindrome1(r)
	}
}

func Benchmark_longestPalindrome2(b *testing.B) {
	r := result()
	for i := 0; i < b.N; i++ {
		longestPalindrome2(r)
	}
}
