package main

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	for _, str := range strs {
		shortest := Min(len(strs[0]), len(str))
		strs[0] = strs[0][0:shortest]
		for index := 0; index < shortest; index++ {
			if strs[0][index] != str[index] {
				strs[0] = strs[0][0:index]
				break
			}
		}
		if strs[0] == "" {
			break
		}
	}
	return strs[0]
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
