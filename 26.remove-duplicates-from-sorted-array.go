package main

/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	count := 0

	for index, num := range nums {
		if index == 0 {
			count++
			continue
		}
		if num != nums[count-1] {
			nums[count] = num
			count++
		}
	}
	return count
}

// @lc code=end
