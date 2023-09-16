package leetcode

// method 1
// 1) use two for loop
// 2) The first loop will be used to scan each number in the target minus nums, while the second loop will confirm whether each result is equal to the following number. If they are equal, we have found two targets.
// TC = O(N^2), SC = (O)1
func twoSum1(nums []int, target int) []int {
    for i := 0; i < len(nums) - 1; i++ {
        result := target - nums[i]

        for j := i + 1; j < len(nums); j++ {
           if result == nums[j] {
               return []int{ i, j }
           }
        }
    }

    return []int{}
}

// method 2
// 1) use one for loop and one map
// 2) when running the for loop, use one resultMap (map[int]int, key is "result", value is "index") to store the result from target minus each num in nums
// 3) during the for loop scanning process, check each "result" and make sure whether it have been store in the map, if it was found in the map that mean we found the two targets
// TC = O(N), SC = O(N)
func twoSum2(nums []int, target int) []int {
    resultMap := map[int]int{}

    for i := 0; i < len(nums); i++ {
        index, ok := resultMap[nums[i]]

        if ok {
            return []int{ index, i }
        }

        resultMap[target - nums[i]] = i
    }

    return []int{}
}