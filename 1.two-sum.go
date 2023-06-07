package main

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	cache := make(map[int]int, 0)
	for index, num := range nums {
		if val, ok := cache[num]; ok {
			return []int{val, index}
		}
		diff := target - num
		cache[diff] = index
	}
	return []int{}
}

// @lc code=end
