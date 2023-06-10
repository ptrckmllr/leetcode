package main

import "sort"

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 *
 * https://leetcode.com/problems/3sum/description/
 *
 * algorithms
 * Medium (32.73%)
 * Likes:    26078
 * Dislikes: 2338
 * Total Accepted:    2.7M
 * Total Submissions: 8.2M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * Given an integer array nums, return all the triplets [nums[i], nums[j],
 * nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] +
 * nums[k] == 0.
 *
 * Notice that the solution set must not contain duplicate triplets.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [-1,0,1,2,-1,-4]
 * Output: [[-1,-1,2],[-1,0,1]]
 * Explanation:
 * nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
 * nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
 * nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
 * The distinct triplets are [-1,0,1] and [-1,-1,2].
 * Notice that the order of the output and the order of the triplets does not
 * matter.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [0,1,1]
 * Output: []
 * Explanation: The only possible triplet does not sum up to 0.
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [0,0,0]
 * Output: [[0,0,0]]
 * Explanation: The only possible triplet sums up to 0.
 *
 *
 *
 * Constraints:
 *
 *
 * 3 <= nums.length <= 3000
 * -10^5 <= nums[i] <= 10^5
 *
 *
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	numsLen := len(nums)
	for i := 0; i < numsLen-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// the key to solve this problem in time was to reduce the
		// problem space from left AND right (my initial approach only went left-to-right)
		right := numsLen - 1
		left := i + 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				//skip duplicates on both ends
				for left+1 < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right-1 && nums[right] == nums[right-1] {
					right--
				}
				// go to next check
				left++
				right--
			}
		}
	}
	return result
}

// @lc code=end

func tooSlow_threeSum(nums []int) [][]int {
	// TODO: too slow, times out
	sort.Ints(nums)
	result := [][]int{}
	cache := map[int]map[int]int{}
	l := len(nums)
	for i := 0; i < l-2; i++ {
		if _, ok := cache[nums[i]]; ok {
			continue
		}
		for j := i + 1; j < l-1; j++ {
			if _, ok := cache[nums[i]][nums[j]]; ok {
				continue
			}
			for k := j + 1; k < l; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					if cache[nums[i]] == nil {
						cache[nums[i]] = map[int]int{}
					}
					cache[nums[i]][nums[j]] = nums[k]
				}
			}
		}
	}
	for k1, v1 := range cache {
		for k2, v2 := range v1 {
			result = append(result, []int{k1, k2, v2})
		}
	}
	return result
}
