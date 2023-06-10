package main

/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 *
 * https://leetcode.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (44.70%)
 * Likes:    6960
 * Dislikes: 7154
 * Total Accepted:    2M
 * Total Submissions: 4.4M
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * A phrase is a palindrome if, after converting all uppercase letters into
 * lowercase letters and removing all non-alphanumeric characters, it reads the
 * same forward and backward. Alphanumeric characters include letters and
 * numbers.
 *
 * Given a string s, return true if it is a palindrome, or false otherwise.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "A man, a plan, a canal: Panama"
 * Output: true
 * Explanation: "amanaplanacanalpanama" is a palindrome.
 *
 *
 * Example 2:
 *
 *
 * Input: s = "race a car"
 * Output: false
 * Explanation: "raceacar" is not a palindrome.
 *
 *
 * Example 3:
 *
 *
 * Input: s = " "
 * Output: true
 * Explanation: s is an empty string "" after removing non-alphanumeric
 * characters.
 * Since an empty string reads the same forward and backward, it is a
 * palindrome.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 2 * 10^5
 * s consists only of printable ASCII characters.
 *
 *
 */

// @lc code=start
func isPalindrome(s string) bool {
	l := len(s)
	if l == 0 || l == 1 {
		return true
	}
	stack := []byte{}
	for i := 0; i < l; i++ {
		c := s[i]
		if c >= 65 && c <= 90 {
			stack = append(stack, c+32)
		} else if (c >= 97 && c <= 122) || (c >= 48 && c <= 57) {
			stack = append(stack, c)
		}
	}
	ls := len(stack)
	if ls == 0 || ls == 1 {
		return true
	}
	for i := 0; i <= ls/2; i++ {
		if stack[i] != stack[ls-1-i] {
			return false
		}
	}
	return true
}

// @lc code=end
