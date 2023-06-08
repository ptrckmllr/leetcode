package main

/*
 * @lc app=leetcode id=2011 lang=golang
 *
 * [2011] Final Value of Variable After Performing Operations
 *
 * https://leetcode.com/problems/final-value-of-variable-after-performing-operations/description/
 *
 * algorithms
 * Easy (88.64%)
 * Likes:    1231
 * Dislikes: 156
 * Total Accepted:    238.5K
 * Total Submissions: 269.1K
 * Testcase Example:  '["--X","X++","X++"]'
 *
 * There is a programming language with only four operations and one variable
 * X:
 *
 *
 * ++X and X++ increments the value of the variable X by 1.
 * --X and X-- decrements the value of the variable X by 1.
 *
 *
 * Initially, the value of X is 0.
 *
 * Given an array of strings operations containing a list of operations, return
 * the final value of X after performing all the operations.
 *
 *
 * Example 1:
 *
 *
 * Input: operations = ["--X","X++","X++"]
 * Output: 1
 * Explanation: The operations are performed as follows:
 * Initially, X = 0.
 * --X: X is decremented by 1, X =  0 - 1 = -1.
 * X++: X is incremented by 1, X = -1 + 1 =  0.
 * X++: X is incremented by 1, X =  0 + 1 =  1.
 *
 *
 * Example 2:
 *
 *
 * Input: operations = ["++X","++X","X++"]
 * Output: 3
 * Explanation: The operations are performed as follows:
 * Initially, X = 0.
 * ++X: X is incremented by 1, X = 0 + 1 = 1.
 * ++X: X is incremented by 1, X = 1 + 1 = 2.
 * X++: X is incremented by 1, X = 2 + 1 = 3.
 *
 *
 * Example 3:
 *
 *
 * Input: operations = ["X++","++X","--X","X--"]
 * Output: 0
 * Explanation: The operations are performed as follows:
 * Initially, X = 0.
 * X++: X is incremented by 1, X = 0 + 1 = 1.
 * ++X: X is incremented by 1, X = 1 + 1 = 2.
 * --X: X is decremented by 1, X = 2 - 1 = 1.
 * X--: X is decremented by 1, X = 1 - 1 = 0.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= operations.length <= 100
 * operations[i] will be either "++X", "X++", "--X", or "X--".
 *
 *
 */

// @lc code=start
func finalValueAfterOperations(operations []string) int {
	value := 0
	for _, s := range operations {
		if s[1] == '+' {
			value++
		} else {
			value--
		}
	}
	return value
}

// @lc code=end
