package main

/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 *
 * https://leetcode.com/problems/product-of-array-except-self/description/
 *
 * algorithms
 * Medium (65.10%)
 * Likes:    17897
 * Dislikes: 985
 * Total Accepted:    1.7M
 * Total Submissions: 2.6M
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given an integer array nums, return an array answer such that answer[i] is
 * equal to the product of all the elements of nums except nums[i].
 *
 * The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
 * integer.
 *
 * You must write an algorithm that runs in O(n) time and without using the
 * division operation.
 *
 *
 * Example 1:
 * Input: nums = [1,2,3,4]
 * Output: [24,12,8,6]
 * Example 2:
 * Input: nums = [-1,1,0,-3,3]
 * Output: [0,0,9,0,0]
 *
 *
 * Constraints:
 *
 *
 * 2 <= nums.length <= 10^5
 * -30 <= nums[i] <= 30
 * The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
 * integer.
 *
 *
 *
 * Follow up: Can you solve the problem in O(1) extra space complexity? (The
 * output array does not count as extra space for space complexity analysis.)
 *
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	product := 1
	zeroes := 0
	for _, v := range nums {
		if v == 0 {
			zeroes++
		} else {
			product *= v
		}
	}
	result := make([]int, len(nums))
	for i := range nums {
		if zeroes > 1 {
			result[i] = 0
		} else if zeroes == 1 {
			if nums[i] == 0 {
				result[i] = product
			} else {
				result[i] = 0
			}
		} else {
			result[i] = product / nums[i]
		}
	}
	return result
}

// @lc code=end
