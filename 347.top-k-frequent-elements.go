package main

import "sort"

/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 *
 * https://leetcode.com/problems/top-k-frequent-elements/description/
 *
 * algorithms
 * Medium (64.05%)
 * Likes:    14631
 * Dislikes: 519
 * Total Accepted:    1.5M
 * Total Submissions: 2.4M
 * Testcase Example:  '[1,1,1,2,2,3]\n2'
 *
 * Given an integer array nums and an integer k, return the k most frequent
 * elements. You may return the answer in any order.
 *
 *
 * Example 1:
 * Input: nums = [1,1,1,2,2,3], k = 2
 * Output: [1,2]
 * Example 2:
 * Input: nums = [1], k = 1
 * Output: [1]
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 * k is in the range [1, the number of unique elements in the array].
 * It is guaranteed that the answer is unique.
 *
 *
 *
 * Follow up: Your algorithm's time complexity must be better than O(n log n),
 * where n is the array's size.
 *
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	cache := map[int]int{}
	for _, num := range nums {
		cache[num]++
	}
	keys := make([]int, 0, len(cache))
	for key := range cache {
		keys = append(keys, key)
	}
	// sort the keys by their map value
	sort.SliceStable(keys, func(i, j int) bool {
		return cache[keys[i]] > cache[keys[j]]
	})
	return keys[:k]
}

// @lc code=end
