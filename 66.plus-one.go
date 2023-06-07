/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {
	remainder := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = digits[i] + remainder
		remainder = digits[i] / 10
		digits[i] = digits[i] % 10
	}
	if remainder != 0 {
		return append([]int{remainder}, digits...)
	}
	return digits
}

// @lc code=end

