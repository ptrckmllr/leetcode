package main

import "log"

/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */

// @lc code=start
func addBinary(a string, b string) string {
	l_a := len(a)
	l_b := len(b)
	l_new := Max(l_a, l_b) + 1
	newStr := make([]byte, l_new)
	i := 0
	remainder := 0
	for {
		current := remainder
		remainder = 0
		i_a := l_a - 1 - i
		if i_a >= 0 && a[i_a] == '1' {
			current++
		}
		i_b := l_b - 1 - i
		if i_b >= 0 && b[i_b] == '1' {
			current++
		}
		i_new := l_new - 1 - i
		if current == 0 {
			newStr[i_new] = '0'
		} else if current == 1 {
			newStr[i_new] = '1'
		} else if current == 2 {
			newStr[i_new] = '0'
			remainder = 1
		} else if current == 3 {
			newStr[i_new] = '1'
			remainder = 1
		}
		i++
		if remainder == 0 && i_a <= 0 && i_b <= 0 {
			break
		}
	}

	return string(newStr[(l_new - i):])
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
type test67 struct {
	a        string
	b        string
	expected string
}

func main() {
	var tests = []test67{
		{
			"11",
			"1",
			"100",
		},
	}
	for _, test := range tests {
		output := addBinary(test.a, test.b)
		log.Printf("[%d = %d]\n", output, test.expected)
	}
}
