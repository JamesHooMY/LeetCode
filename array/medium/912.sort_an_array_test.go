package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sort-an-array/description/

// method 1 selection sort, small numbers keep in the left
// 1) iterate the array, always find the min number in the rest of the array
// 2) swap the min number with the initial number of each iteration
// TC: O(N^2), SC: O(1)
func sortArray1(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}

	for i := 0; i < n; i++ {
		minIndex := i

		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}

		if minIndex != i {
			nums[i], nums[minIndex] = nums[minIndex], nums[i]
		}
	}

	return nums
}

// method 2 bubble sort, big numbers keep in the right
// 1) iterate the array, always compare the current number with the next number
// 2) swap the current number with the next number if the current number is bigger than the next number
// TC: O(N^2), SC: O(1)
func sortArray2(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}

// method 3 insertion sort, this is the best performance with TC: O(N^2), which can pass the leetcode test
// 1) iterate the array, always compare the current number with the previous numbers
// 2) if the previous number is bigger than the current number, move the previous number to the next position
// 3) insert the current number into the preIndex+1 position
// TC: O(N^2), SC: O(1)
func sortArray3(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	for i := 1; i < len(nums); i++ {
		preIndex := i - 1
		current := nums[i]

		for preIndex >= 0 && nums[preIndex] > current {
			nums[preIndex+1] = nums[preIndex]
			preIndex--
		}

		// insert the current number into the preIndex+1 position, due to preIndex-- in the for loop
		nums[preIndex+1] = current
	}

	return nums
}

// method 4 quick sort (recursion + two pointers in-place partitioning), https://www.youtube.com/watch?v=Hoixgm4-P4M&list=RDCMUCzDJwLWoYCUQowF_nG3m5OQ&index=1
// 1) select a pivot from middle of the array, move the pivot to the first element of the array
// 2) use two pointers to operate the in-place partitioning(optimized for reducing memory usage), separate the array into two parts, left part is smaller than the pivot, right part is bigger than the pivot
// 3) swap the pivot with the right pointer
// 4) recursion the left part and right part
// TC: O(NlogN), SC: O(logN)
func sortArray4(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	left, right := 0, len(nums)-1
	pivotIndex := partition(nums, left, right)

	sortArray4(nums[:pivotIndex])
	sortArray4(nums[pivotIndex+1:])

	return nums
}

// partition function separates the array into two parts, left parts are smaller than the pivot, right parts are bigger or equal to the pivot
func partition(nums []int, left, right int) int {
	// * always select the first element as the pivot -----------------------
	// pivot := nums[left]
	// rightIdx := right

	// * keep the elements after rightIdx are bigger than the pivot
	// for i := right; i > left; i-- {
	// 	if nums[i] > pivot {
	// 		nums[i], nums[rightIdx] = nums[rightIdx], nums[i]
	// 		rightIdx--
	// 	}
	// }

	// nums[left], nums[rightIdx] = nums[rightIdx], nums[left]

	// return rightIdx
	// * always select the first element as the pivot -----------------------

	// * always select the last element as the pivot -----------------------
	// pivot := nums[right]
	// leftIdx := left

	// * keep the elements before leftIdx are smaller than the pivot
	// for i := left; i < right; i++ {
	// 	if nums[i] < pivot {
	// 		nums[i], nums[leftIdx] = nums[leftIdx], nums[i]
	// 		leftIdx++
	// 	}
	// }

	// nums[leftIdx], nums[right] = nums[right], nums[leftIdx]

	// return leftIdx
	// * always select the last element as the pivot -----------------------

	// * select the middle element as the pivot, pivotIndex := rand.Intn(len(nums)) is better for reducing the worst case, but it will increase the time complexity
	pivotIndex := len(nums) / 2
	pivot := nums[pivotIndex]

	// * pivot was set to the first element ------------------------------------------------------
	// nums[pivotIndex], nums[0] = nums[0], nums[pivotIndex]

	// for left <= right {
	// 	for left <= right && nums[left] <= pivot {
	// 		left++
	// 	}

	// 	// * nums[right] > pivot, make the elements after right pointer are bigger than the pivot
	// 	for left <= right && nums[right] > pivot {
	// 		right--
	// 	}

	// 	if left < right {
	// 		nums[left], nums[right] = nums[right], nums[left]
	// 	}
	// }

	// // swap the pivot to the right pointer
	// nums[0], nums[right] = nums[right], nums[0]

	// return right
	// * pivot was set to the first element ------------------------------------------------------

	// * pivot was set to the last element ------------------------------------------------------
	nums[pivotIndex], nums[len(nums)-1] = nums[len(nums)-1], nums[pivotIndex]

	for left <= right {
		// * nums[left] < pivot, make the elements before left pointer are smaller than the pivot
		for left <= right && nums[left] < pivot {
			left++
		}

		for left <= right && nums[right] >= pivot {
			right--
		}

		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}

	// swap the pivot to the left pointer
	nums[len(nums)-1], nums[left] = nums[left], nums[len(nums)-1]

	return left
	// * pivot was set to the last element ------------------------------------------------------
}

// method 5 merge sort (recursion + divide and conquer) top-down, easy to understand, https://www.youtube.com/watch?v=4VqmGXwpLqc&list=RDCMUCzDJwLWoYCUQowF_nG3m5OQ&index=2
// 1) divide the array into two parts, mid = len(nums) / 2
// 2) recursion the left part (nums[:mid]) and right part (nums[mid:])
// 3) merge the left part and right part
// TC: O(NlogN), SC: O(N)
func sortArray5(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2

	leftArr := sortArray5(nums[:mid])
	rightArr := sortArray5(nums[mid:])

	/*
		example: [5, 2, 3, 1], len(nums) => 4

		leftArr => [5, 2], rightArr => [3, 1]
		* [5, 2] separated into [5] and [2], then merge [5] and [2] into [2, 5]
		* [3, 1] separated into [3] and [1], then merge [3] and [1] into [1, 3]
	*/

	return mergeSort(leftArr, rightArr)
}

func mergeSort(leftArr, rightArr []int) []int {
	result := make([]int, 0, len(leftArr)+len(rightArr))

	leftIdx, rightIdx := 0, 0
	for leftIdx < len(leftArr) && rightIdx < len(rightArr) {
		if leftArr[leftIdx] <= rightArr[rightIdx] {
			result = append(result, leftArr[leftIdx])
			leftIdx++
		} else {
			result = append(result, rightArr[rightIdx])
			rightIdx++
		}
	}

	// if leftIdx < len(leftArr) {
	// 	result = append(result, leftArr[leftIdx:]...)
	// } else {
	// 	result = append(result, rightArr[rightIdx:]...)
	// }
	result = append(result, leftArr[leftIdx:]...)
	result = append(result, rightArr[rightIdx:]...)

	/*
		slice operation example
		s := []int{1}

		s1 := s[len(s):] => [], this will not out of range, the result is empty slice
		s1 := s[len(s)+1:] => out of range
		s1 := s[len(s)-1:] => [1]
	*/

	return result
}

// method 6 merge sort (iteration + divide and conquer) bottom-up, https://www.youtube.com/watch?v=IN_ZOU-LK08
// 1) iterate the array, always merge two parts into one part
// 2) step *= 2, i += 2 * step
// 3) left = i is the start index of the left part, mid = i + step is the start index of the right part, right = i + 2 * step is the end index of the right part
// 4) if right > n, right = n, due to the right part may not have 2 * step elements
// 5) merge the left part and right part
// TC: O(NlogN), SC: O(N)
func sortArray6(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}

	/*
		example: [5, 2, 3, 1, 6, 4], len(nums) => n = 6
		step => 1 2 4
	*/
	for step := 1; step < n; step *= 2 {
		/*
			step => 1, n - step => 5, i => 0 2 4
			step => 2, n - step => 4, i => 0
			step => 4, n - step => 2, i => 0
		*/
		for i := 0; i < n-step; i += 2 * step {
			left := i
			mid := i + step

			right := i + 2*step
			if right > n {
				right = n
			}
			/*
				step => 1:
					i => 0:
						left => 0, mid => 1, right => 2
						[5(left), 2(mid), 3(right), 1, 6, 4] => [2, 5, 3, 1, 6, 4]
						* [5, 2] separated into [5] and [2], then merge [5] and [2] into [2, 5]

					i => 2:
						left => 2, mid => 3, right => 4
						[2, 5, 3(left), 1(mid), 6(right), 4] => [2, 5, 1, 3, 6, 4]
						* [3, 1] separated into [3] and [1], then merge [3] and [1] into [1, 3]

					i => 4:
						left => 4, mid => 5, right => 6
						[2, 5, 1, 3, 6(left), 4(mid)] => [2, 5, 1, 3, 4, 6]
						* [6, 4] separated into [6] and [4], then merge [6] and [4] into [4, 6]

				step => 2:
					i => 0:
						left => 0, mid => 2, right => 4
						[2(left), 5, 1(mid), 3, 4(right), 6] => [1, 2, 5, 3, 4, 6]
						* [2, 5, 1, 3] separated into [2, 5] and [1, 3], then merge [2, 5] and [1, 3] into [1, 2, 5, 3]

				step => 4:
					i => 0:
						left => 0, mid => 4, right => 6 (right != 8, due to right > n, right = n)
						[1(left), 2, 5, 3, 4(mid), 6] => [1, 2, 3, 4, 5, 6]
						* [1, 2, 5, 3, 4, 6] separated into [1, 2, 5, 3] and [4, 6], then merge [1, 2, 5, 3] and [4, 6] into [1, 2, 3, 4, 5, 6]
			*/

			leftArr := nums[left:mid]
			rightArr := nums[mid:right]

			merged := mergeSort(leftArr, rightArr)
			// copy(nums[left:left+len(merged)], merged)
			copy(nums[left:right], merged)
		}
	}

	return nums
}

// method 7 heap sort
// 1) build the max heap with maxHeapify functionn, n/2-1 is the index of the last non-leaf node
// 2) extract the max element from the heap, then swap the max element with the last element of the heap
// 3) heapify the heap from the root node, the length of the heap is decreased by 1 after each iteration
// 4) finally, the array is sorted
// TC: O(NlogN), SC: O(1)
// * this is the best solution for me currently
func sortArray7(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}

	/*
		build the max heap, n/2-1 is the index of the last non-leaf node, bottom-up

		example: [5, 2, 3, 1, 6, 4], len(nums) => n = 6
		* n/2-1 => 2, i => 2 1 0

		i => 2:
		heapify(nums, 6, 2) => 3 <-> 4 => [5, 2, 4, 1, 6, 3]

		i => 1:
		heapify(nums, 6, 1) => 2 <-> 6 => [5, 6, 4, 1, 2, 3]

		i => 0:
		heapify(nums, 6, 0) => 5 <-> 6 => [6, 5, 4, 1, 2, 3]
	*/
	for i := n/2 - 1; i >= 0; i-- { // TC: O(N), SC: O(1)
		maxHeapify(nums, n, i) // TC: O(logN), SC: O(1)
	}

	/*
		change the max element with the last element, then heapify the heap from the root node, top-down

		example: [6, 5, 4, 1, 2, 3], len(nums) => n = 6
		* i => 5 4 3 2 1 0

		i => 5:
		nums[0], nums[5] = nums[5], nums[0] => [3, 5, 4, 1, 2, 6]
		heapify(nums, 5, 0) => 3 <-> 5 => [5, 3, 4, 1, 2, 6]

		i => 4:
		nums[0], nums[4] = nums[4], nums[0] => [2, 3, 4, 1, 5, 6]
		heapify(nums, 4, 0) => 2 <-> 4 => [4, 3, 2, 1, 5, 6]

		i => 3:
		nums[0], nums[3] = nums[3], nums[0] => [1, 3, 2, 4, 5, 6]
		heapify(nums, 3, 0) => 1 <-> 3 => [3, 1, 2, 4, 5, 6]

		i => 2:
		nums[0], nums[2] = nums[2], nums[0] => [2, 1, 3, 4, 5, 6]
		heapify(nums, 2, 0) => [2, 1, 3, 4, 5, 6]

		i => 1:
		nums[0], nums[1] = nums[1], nums[0] => [1, 2, 3, 4, 5, 6]
		heapify(nums, 1, 0) => [1, 2, 3, 4, 5, 6]
	*/
	for i := n - 1; i >= 0; i-- { // TC: O(N), SC: O(1)
		nums[0], nums[i] = nums[i], nums[0]
		maxHeapify(nums, i, 0) // TC: O(logN), SC: O(1)
	}

	return nums
}

// TC: O(logN), SC: O(1)
func maxHeapify(nums []int, n, i int) {
	// i is the index of the node for heapify start from
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	// find the largest element among the parent node and its children
	if left < n && nums[left] > nums[largest] {
		largest = left
	}

	if right < n && nums[right] > nums[largest] {
		largest = right
	}

	// if the largest element is not the parent node, swap the largest element with the parent node
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		maxHeapify(nums, n, largest)
	}
}

// selection sort
func Test_sortArray1(t *testing.T) {
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
				nums: []int{5, 2, 3, 1},
			},
			expected: expected{
				result: []int{1, 2, 3, 5},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{5, 1, 1, 2, 0, 0},
			},
			expected: expected{
				result: []int{0, 0, 1, 1, 2, 5},
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{5, 1, 1, 2, 0, 0, 3, 4, 6, 7, 8, 9, 10, 11, 12, 13},
			},
			expected: expected{
				result: []int{0, 0, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
			},
		},
		{
			name: "4",
			args: args{
				nums: []int{-4, 0, 7, 4, 9, -5, -1, 0, -7, -1},
			},
			expected: expected{
				result: []int{-7, -5, -4, -1, -1, 0, 0, 4, 7, 9},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			sortArray1(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// bubble sort
func Test_sortArray2(t *testing.T) {
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
				nums: []int{5, 2, 3, 1},
			},
			expected: expected{
				result: []int{1, 2, 3, 5},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{5, 1, 1, 2, 0, 0},
			},
			expected: expected{
				result: []int{0, 0, 1, 1, 2, 5},
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{5, 1, 1, 2, 0, 0, 3, 4, 6, 7, 8, 9, 10, 11, 12, 13},
			},
			expected: expected{
				result: []int{0, 0, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
			},
		},
		{
			name: "4",
			args: args{
				nums: []int{-4, 0, 7, 4, 9, -5, -1, 0, -7, -1},
			},
			expected: expected{
				result: []int{-7, -5, -4, -1, -1, 0, 0, 4, 7, 9},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			sortArray2(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// insertion sort
func Test_sortArray3(t *testing.T) {
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
				nums: []int{5, 2, 3, 1},
			},
			expected: expected{
				result: []int{1, 2, 3, 5},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{5, 1, 1, 2, 0, 0},
			},
			expected: expected{
				result: []int{0, 0, 1, 1, 2, 5},
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{5, 1, 1, 2, 0, 0, 3, 4, 6, 7, 8, 9, 10, 11, 12, 13},
			},
			expected: expected{
				result: []int{0, 0, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
			},
		},
		{
			name: "4",
			args: args{
				nums: []int{-4, 0, 7, 4, 9, -5, -1, 0, -7, -1},
			},
			expected: expected{
				result: []int{-7, -5, -4, -1, -1, 0, 0, 4, 7, 9},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			sortArray3(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// quick sort
func Test_sortArray4(t *testing.T) {
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
				nums: []int{5, 2, 3, 1},
			},
			expected: expected{
				result: []int{1, 2, 3, 5},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{5, 1, 1, 2, 0, 0},
			},
			expected: expected{
				result: []int{0, 0, 1, 1, 2, 5},
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{5, 1, 1, 2, 0, 0, 3, 4, 6, 7, 8, 9, 10, 11, 12, 13},
			},
			expected: expected{
				result: []int{0, 0, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
			},
		},
		{
			name: "4",
			args: args{
				nums: []int{-4, 0, 7, 4, 9, -5, -1, 0, -7, -1},
			},
			expected: expected{
				result: []int{-7, -5, -4, -1, -1, 0, 0, 4, 7, 9},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			sortArray4(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// merge sort (recursion + divide and conquer) top-down
func Test_sortArray5(t *testing.T) {
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
				nums: []int{5, 2, 3, 1},
			},
			expected: expected{
				result: []int{1, 2, 3, 5},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{5, 1, 1, 2, 0, 0},
			},
			expected: expected{
				result: []int{0, 0, 1, 1, 2, 5},
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{5, 1, 1, 2, 0, 0, 3, 4, 6, 7, 8, 9, 10, 11, 12, 13},
			},
			expected: expected{
				result: []int{0, 0, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
			},
		},
		{
			name: "4",
			args: args{
				nums: []int{-4, 0, 7, 4, 9, -5, -1, 0, -7, -1},
			},
			expected: expected{
				result: []int{-7, -5, -4, -1, -1, 0, 0, 4, 7, 9},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			sortArray5(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// merge sort (iteration + divide and conquer) bottom-up
func Test_sortArray6(t *testing.T) {
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
				nums: []int{5, 2, 3, 1},
			},
			expected: expected{
				result: []int{1, 2, 3, 5},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{5, 1, 1, 2, 0, 0},
			},
			expected: expected{
				result: []int{0, 0, 1, 1, 2, 5},
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{5, 1, 1, 2, 0, 0, 3, 4, 6, 7, 8, 9, 10, 11, 12, 13},
			},
			expected: expected{
				result: []int{0, 0, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
			},
		},
		{
			name: "4",
			args: args{
				nums: []int{-4, 0, 7, 4, 9, -5, -1, 0, -7, -1},
			},
			expected: expected{
				result: []int{-7, -5, -4, -1, -1, 0, 0, 4, 7, 9},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			sortArray6(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// heap sort
func Test_sortArray7(t *testing.T) {
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
				nums: []int{5, 2, 3, 1},
			},
			expected: expected{
				result: []int{1, 2, 3, 5},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{5, 1, 1, 2, 0, 0},
			},
			expected: expected{
				result: []int{0, 0, 1, 1, 2, 5},
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{5, 1, 1, 2, 0, 0, 3, 4, 6, 7, 8, 9, 10, 11, 12, 13},
			},
			expected: expected{
				result: []int{0, 0, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
			},
		},
		{
			name: "4",
			args: args{
				nums: []int{-4, 0, 7, 4, 9, -5, -1, 0, -7, -1},
			},
			expected: expected{
				result: []int{-7, -5, -4, -1, -1, 0, 0, 4, 7, 9},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			sortArray7(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

var nums = []int{5, 1, 1, 2, 0, 0, 3, 4, 6, 7, 8, 9, 10, 11, 12, 13}

// benchmark

// selection sort
func Benchmark_sortArray1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortArray1(nums)
	}
}

// bubble sort
func Benchmark_sortArray2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortArray2(nums)
	}
}

// insertion sort
func Benchmark_sortArray3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortArray3(nums)
	}
}

// quick sort
func Benchmark_sortArray4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortArray4(nums)
	}
}

// merge sort (recursion + divide and conquer) top-down
func Benchmark_sortArray5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortArray5(nums)
	}
}

// merge sort (iteration + divide and conquer) bottom-up
func Benchmark_sortArray6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortArray6(nums)
	}
}

// heap sort
func Benchmark_sortArray7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortArray7(nums)
	}
}
