package main

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var digits [10]int
	var digitCount int
	for index, _ := range digits {
		digits[index] = x % 10
		x = x / 10
		if x == 0 {
			digitCount = index + 1
			break
		}
	}

	for i := 0; i < digitCount/2; i++ {
		if digits[i] != digits[digitCount-1-i] {
			return false
		}
	}
	return true
}

// @lc code=end
