package main

/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	characterValues := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	previous := 0
	for _, character := range s {
		value := characterValues[character]
		total = total + value
		if value > previous {
			total = total - 2*previous
		}
		previous = value
	}
	return total
}

// @lc code=end
