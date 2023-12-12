package easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/valid-anagram/description/

// method 1 use map
// 1) check if the length of s and t are equal, if not, then return false
// 2) use a map to store the characters of s and t
// 3) if the character is in s, then add 1 to the value of the map
// 4) if the character is in t, then subtract 1 to the value of the map
// 5) if the value of the map is not 0, then return false
// 6) finally, return true
// TC = O(N), SC = O(N)
func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charCountMap := make(map[rune]int) // key: character, value: count
	for i := range s {
		charCountMap[rune(s[i])]++
		charCountMap[rune(t[i])]--
	}
	// * byte in map is slower than rune ? howerver in 125.valid_palindrome_test.go, byte is faster than rune in IsAlphanumeric and ToLowerCase. why ?
	// charCountMap := make(map[byte]int)
	// for i := range s {
	// 	charCountMap[byte(s[i])]++
	// 	charCountMap[byte(t[i])]--
	// }

	for _, count := range charCountMap {
		if count != 0 {
			return false
		}
	}

	return true
}

// method 2 use array, manipulate the index of the array is faster than manipulate the value of the map
// 1) check if the length of s and t are equal, if not, then return false
// 2) use an array to store the characters of s and t
// 3) if the character is in s, then add 1 to the value of the array
// 4) if the character is in t, then subtract 1 to the value of the array
// 5) if the value of the array is not 0, then return false
// 6) finally, return true
// TC = O(N), SC = O(1)
// * this is the best solution for me currently
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// make an array with 26 length, the memory address of the array will not be changed during the execution of the program
	charArr := make([]int, 26) // index: character - 'a', value: count
	for i := range s {
		// * this is the key point, use charArr[s[i]-'a']++ instead of charArr[s[i]]++
		// s[i]-'a' and t[i]-'a' are the index of the array, subtracting 'a' is to make the index start from 0, charArr is make with 26 length, so the index is from 0 to 25, if not subtracting 'a', then the index will be from 97 to 122, this will cause index out of range
		charArr[s[i]-'a']++
		charArr[t[i]-'a']--
	}

	for _, count := range charArr {
		if count != 0 {
			return false
		}
	}

	return true
}

func Test_isAnagram1(t *testing.T) {
	type args struct {
		s string
		t string
	}
	type expected struct {
		result bool
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
				s: "anagram",
				t: "nagaram",
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				s: "rat",
				t: "car",
			},
			expected: expected{
				result: false,
			},
		},
		{
			name: "3",
			args: args{
				s: "a",
				t: "ab",
			},
			expected: expected{
				result: false,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			isAnagram1(tc.args.s, tc.args.t),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_isAnagram2(t *testing.T) {
	type args struct {
		s string
		t string
	}
	type expected struct {
		result bool
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
				s: "anagram",
				t: "nagaram",
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				s: "rat",
				t: "car",
			},
			expected: expected{
				result: false,
			},
		},
		{
			name: "3",
			args: args{
				s: "a",
				t: "ab",
			},
			expected: expected{
				result: false,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			isAnagram2(tc.args.s, tc.args.t),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_isAnagram1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isAnagram1("anagram", "nagaram")
	}
}

func Benchmark_isAnagram2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isAnagram2("anagram", "nagaram")
	}
}
