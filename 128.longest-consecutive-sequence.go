package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 *
 * https://leetcode.com/problems/longest-consecutive-sequence/description/
 *
 * algorithms
 * Medium (48.23%)
 * Likes:    16359
 * Dislikes: 690
 * Total Accepted:    1.1M
 * Total Submissions: 2.4M
 * Testcase Example:  '[100,4,200,1,3,2]'
 *
 * Given an unsorted array of integers nums, return the length of the longest
 * consecutive elements sequence.
 *
 * You must write an algorithm that runs in O(n) time.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [100,4,200,1,3,2]
 * Output: 4
 * Explanation: The longest consecutive elements sequence is [1, 2, 3, 4].
 * Therefore its length is 4.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [0,3,7,2,5,8,4,6,0,1]
 * Output: 9
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 *
 *
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	maxChain := 0
	currentChain := 0
	for i, v := range nums {
		if i == 0 {
			currentChain++
			continue
		}
		diff := v - nums[i-1]
		if diff == 1 {
			currentChain++
		} else if diff == 0 {
			continue
		} else {
			if maxChain < currentChain {
				maxChain = currentChain
			}
			currentChain = 1
		}
	}
	if maxChain < currentChain {
		return currentChain
	}
	return maxChain
}

// @lc code=end
