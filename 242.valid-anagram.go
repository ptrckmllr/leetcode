package main

/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 *
 * https://leetcode.com/problems/valid-anagram/description/
 *
 * algorithms
 * Easy (63.08%)
 * Likes:    9277
 * Dislikes: 292
 * Total Accepted:    2.2M
 * Total Submissions: 3.5M
 * Testcase Example:  '"anagram"\n"nagaram"'
 *
 * Given two strings s and t, return true if t is an anagram of s, and false
 * otherwise.
 *
 * An Anagram is a word or phrase formed by rearranging the letters of a
 * different word or phrase, typically using all the original letters exactly
 * once.
 *
 *
 * Example 1:
 * Input: s = "anagram", t = "nagaram"
 * Output: true
 * Example 2:
 * Input: s = "rat", t = "car"
 * Output: false
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length, t.length <= 5 * 10^4
 * s and t consist of lowercase English letters.
 *
 *
 *
 * Follow up: What if the inputs contain Unicode characters? How would you
 * adapt your solution to such a case?
 *
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	len_s := len(s)
	len_t := len(t)
	if len_s != len_t {
		return false
	}
	cache := map[byte]int{}
	for i := 0; i < len_s; i++ {
		cache[s[i]]++
		cache[t[i]]--
	}
	for _, v := range cache {
		if v != 0 {
			return false
		}
	}
	return true
}

// @lc code=end
