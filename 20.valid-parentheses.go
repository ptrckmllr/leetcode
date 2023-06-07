package main

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	stack := []rune{}
	for _, c := range s {
		switch c {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		default:
			l := len(stack)
			if l == 0 {
				return false
			}
			if stack[l-1] != c {
				return false
			}
			stack = stack[:l-1]
		}
	}
	return len(stack) == 0
}

// @lc code=end
