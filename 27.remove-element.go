package main

/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	count := 0
	pointer := len(nums) - 1
	for index, num := range nums {
		if num == val {
			for i := pointer; i >= 0; i-- {
				if nums[i] != val {
					nums[index] = nums[i]
					pointer = i - 1
					break
				}
			}
		} else {
			count++
		}
	}
	return count
}

// @lc code=end
