package arrays

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]
*/
func SolveTwoSum(nums []int, target int) []int {
	// target - current number = number we are looking for
	// does that number exist in the number we iterated so far?
	// since the result is indexes hashmap should be maps[number, index]

	numMap := make(map[int]int, len(nums))
	result := make([]int, 2, 2)

	for i, v := range nums {
		numToLook := target - v
		value, exist := numMap[numToLook]

		if !exist {
			numMap[v] = i
		} else {
			result[0] = i
			result[1] = value
			break
		}
	}

	return result
}
