package main

/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 *
 */
// NOTE: fibonacci
// @lc code=start
func climbStairs(n int) int {
	a, b, c := 1, 1, 0
	for i := 2; i <= n; i++ {
		c = a + b
		b = a
		a = c
	}
	return a
}

// @lc code=end

func main() {
	println(climbStairs(8))
}
