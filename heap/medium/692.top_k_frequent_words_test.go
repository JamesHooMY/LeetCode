package medium

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/top-k-frequent-words/description/

type WordFreq struct {
	word string
	freq int
}

type MinHeapWordFreq []WordFreq

func (h *MinHeapWordFreq) Len() int { return len(*h) }

// min heap
func (h *MinHeapWordFreq) Less(i, j int) bool {
	// this condition is specific for this problem to get the right result
	if (*h)[i].freq == (*h)[j].freq {
		// this order is according to the lexicographical order
		return (*h)[i].word > (*h)[j].word
	}
	return (*h)[i].freq < (*h)[j].freq
}

func (h *MinHeapWordFreq) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MinHeapWordFreq) Push(x any) { *h = append(*h, x.(WordFreq)) }

func (h *MinHeapWordFreq) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// method 1 min heap
// 1) build a map to store the frequency of each word
// 2) build a min heap to store the word and its frequency
// 3) keep the size of heap to k
// 4) pop the top k words from the heap
// TC: O(NlogK), SC: O(N)
// * this is the best solution for me currently
func topKFrequent1(words []string, k int) []string {
	wordFreq := make(map[string]int)
	for _, word := range words {
		wordFreq[word]++
	}

	h := &MinHeapWordFreq{}
	for word, freq := range wordFreq {
		heap.Push(h, WordFreq{word, freq})

		// * this is the key point, keep the heap size to k
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).(WordFreq).word
	}

	return result
}

type MaxHeapWordFreq []WordFreq

func (h *MaxHeapWordFreq) Len() int { return len(*h) }

func (h *MaxHeapWordFreq) Less(i, j int) bool {
	if (*h)[i].freq == (*h)[j].freq {
		return (*h)[i].word < (*h)[j].word
	}

	return (*h)[i].freq > (*h)[j].freq
}

func (h *MaxHeapWordFreq) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MaxHeapWordFreq) Push(x any) { *h = append(*h, x.(WordFreq)) }

func (h *MaxHeapWordFreq) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// method 2 max heap (this method is slower than method 1)
// 1) build a map to store the frequency of each word
// 2) build a max heap to store the word and its frequency
// 3) pop the top k words from the heap
// TC: O(NlogN), SC: O(N)
func topKFrequent2(words []string, k int) []string {
	wordFreq := make(map[string]int)
	for _, word := range words {
		wordFreq[word]++
	}

	h := &MaxHeapWordFreq{}
	for word, freq := range wordFreq {
		// if the length of wordFreq is large, heap.Push() will be slower due to the comparison with parent node too many times
		heap.Push(h, WordFreq{word, freq})
	}

	result := make([]string, k)
	for i := 0; i < k; i++ {
		result[i] = heap.Pop(h).(WordFreq).word
	}

	return result
}

// method 3 bucket sort
// 1) build a map to store the frequency of each word
// 2) build a bucket to store the words with the same frequency
// 3) iterate the buckets from the end(highest frequency) to the beginning(lowest frequency)
// 4) sort the words in the buckets by lexicographical order
// 5) append the words to result
// TC: O(NlogN), SC: O(N)
func topKFrequent3(words []string, k int) []string {
	wordFreqMap := make(map[string]int)
	for _, word := range words {
		wordFreqMap[word]++
	}

	// put the words with the same frequency into the same bucket
	buckets := make([][]string, len(words)+1) // index: frequency of the word
	for word, freq := range wordFreqMap {
		buckets[freq] = append(buckets[freq], word)
	}

	result := make([]string, 0, k) // length: 0, capacity: k
	// iterate the buckets from the end(highest frequency) to the beginning(lowest frequency)
	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		if len(buckets[i]) == 0 {
			continue
		}

		// sort the words in the buckets by lexicographical order
		// TC: O(NlogN), SC: O(N)
		sort.Strings(buckets[i])
		// sort the words with the same frequency by lexicographical order
		// sort.Slice(buckets[i], func(j, l int) bool {
		// 	return buckets[i][j] < buckets[i][l]
		// })

		for _, word := range buckets[i] {
			result = append(result, word)
			if len(result) == k {
				break
			}
		}
		// check len(buckets[i]) to avoid out of range error when append the words to result
		// if len(buckets[i]) > k-len(result) {
		// 	result = append(result, buckets[i][:k-len(result)]...)
		// } else {
		// 	result = append(result, buckets[i]...)
		// }
	}

	return result
}

func Test_topKFrequent1(t *testing.T) {
	type args struct {
		words []string
		k     int
	}
	type expected struct {
		result []string
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
				words: []string{"i", "love", "leetcode", "i", "love", "coding"},
				k:     2,
			},
			expected: expected{
				result: []string{"i", "love"},
			},
		},
		{
			name: "2",
			args: args{
				words: []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"},
				k:     4,
			},
			expected: expected{
				result: []string{"the", "is", "sunny", "day"},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			topKFrequent1(tc.args.words, tc.args.k),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_topKFrequent2(t *testing.T) {
	type args struct {
		words []string
		k     int
	}
	type expected struct {
		result []string
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
				words: []string{"i", "love", "leetcode", "i", "love", "coding"},
				k:     2,
			},
			expected: expected{
				result: []string{"i", "love"},
			},
		},
		{
			name: "2",
			args: args{
				words: []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"},
				k:     4,
			},
			expected: expected{
				result: []string{"the", "is", "sunny", "day"},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			topKFrequent2(tc.args.words, tc.args.k),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_topKFrequent3(t *testing.T) {
	type args struct {
		words []string
		k     int
	}
	type expected struct {
		result []string
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
				words: []string{"i", "love", "leetcode", "i", "love", "coding"},
				k:     2,
			},
			expected: expected{
				result: []string{"i", "love"},
			},
		},
		{
			name: "2",
			args: args{
				words: []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"},
				k:     4,
			},
			expected: expected{
				result: []string{"the", "is", "sunny", "day"},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			topKFrequent3(tc.args.words, tc.args.k),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_topKFrequent1(b *testing.B) {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	k := 2
	for i := 0; i < b.N; i++ {
		topKFrequent1(words, k)
	}
}

func Benchmark_topKFrequent2(b *testing.B) {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	k := 2
	for i := 0; i < b.N; i++ {
		topKFrequent2(words, k)
	}
}

func Benchmark_topKFrequent3(b *testing.B) {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	k := 2
	for i := 0; i < b.N; i++ {
		topKFrequent3(words, k)
	}
}
