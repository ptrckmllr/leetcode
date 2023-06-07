package main

/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Find the Index of the First Occurrence in a String
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	needleLen := len(needle)
	needleIndex := 0
	for index := 0; index < len(haystack); index++ {
		if haystack[index] == needle[needleIndex] {
			if needleIndex == needleLen-1 {
				return index - needleIndex
			}
			needleIndex++
			continue
		}
		index = index - needleIndex
		needleIndex = 0
	}
	return -1
}

// @lc code=end
