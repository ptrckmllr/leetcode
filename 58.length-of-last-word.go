package main

/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != 32 {
			length++
		}
		if s[i] == 32 && length != 0 {
			return length
		}
	}
	return length
}

// @lc code=end
