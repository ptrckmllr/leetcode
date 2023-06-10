package main

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 *
 * https://leetcode.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (53.98%)
 * Likes:    24538
 * Dislikes: 1312
 * Total Accepted:    2.1M
 * Total Submissions: 4M
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * You are given an integer array height of length n. There are n vertical
 * lines drawn such that the two endpoints of the i^th line are (i, 0) and (i,
 * height[i]).
 *
 * Find two lines that together with the x-axis form a container, such that the
 * container contains the most water.
 *
 * Return the maximum amount of water a container can store.
 *
 * Notice that you may not slant the container.
 *
 *
 * Example 1:
 *
 *
 * Input: height = [1,8,6,2,5,4,8,3,7]
 * Output: 49
 * Explanation: The above vertical lines are represented by array
 * [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the
 * container can contain is 49.
 *
 *
 * Example 2:
 *
 *
 * Input: height = [1,1]
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * n == height.length
 * 2 <= n <= 10^5
 * 0 <= height[i] <= 10^4
 *
 *
 */

// @lc code=start
func maxArea(heights []int) int {
	// TODO: the fast solution does the same as in problem 15.3-sum
	// need to reduce problem space from left and right simultaneously
}

// @lc code=end
func slow_maxArea(heights []int) int {
	maxArea := 0
	highestLeft := 0
	for i := 0; i < len(heights)-1; i++ {
		if heights[i] < highestLeft {
			continue
		}
		highestLeft = heights[i]
		highestRight := 0
		for j := len(heights) - 1; j > i; j-- {
			if heights[j] < highestRight {
				continue
			}
			highestRight = heights[j]
			var height int
			var width int
			width = j - i
			if heights[i] > heights[j] {
				height = heights[j]
			} else {
				height = heights[i]
			}
			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}
