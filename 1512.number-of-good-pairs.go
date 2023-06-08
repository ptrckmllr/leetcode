package main

/*
 * @lc app=leetcode id=1512 lang=golang
 *
 * [1512] Number of Good Pairs
 *
 * https://leetcode.com/problems/number-of-good-pairs/description/
 *
 * algorithms
 * Easy (88.21%)
 * Likes:    3932
 * Dislikes: 193
 * Total Accepted:    451.1K
 * Total Submissions: 511.4K
 * Testcase Example:  '[1,2,3,1,1,3]'
 *
 * Given an array of integers nums, return the number of good pairs.
 *
 * A pair (i, j) is called good if nums[i] == nums[j] and i < j.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,2,3,1,1,3]
 * Output: 4
 * Explanation: There are 4 good pairs (0,3), (0,4), (3,4), (2,5) 0-indexed.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1,1,1,1]
 * Output: 6
 * Explanation: Each pair in the array are good.
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [1,2,3]
 * Output: 0
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 100
 * 1 <= nums[i] <= 100
 *
 *
 */

// @lc code=start
func numIdenticalPairs(nums []int) int {
	cache := map[int][]int{}
	for index, num := range nums {
		if cache[num] == nil {
			cache[num] = []int{}
		}
		cache[num] = append(cache[num], index)
	}
	value := 0
	for _, v := range cache {
		value += goodPairs(len(v))
	}
	return value
}
func goodPairs(n int) int {
	if n == 1 {
		return 0
	}
	return n - 1 + goodPairs(n-1)
}

// @lc code=end
