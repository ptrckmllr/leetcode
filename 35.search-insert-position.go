package main

import (
	"log"
)

/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */
// NOTE: never implement binary search yourself
// @lc code=start

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)

	for low < high {
		mid := (high + low) / 2
		if target < nums[mid] {
			high = mid
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return low
}

// @lc code=end
type test struct {
	nums     []int
	target   int
	expected int
}

func main() {
	var tests = []test{
		{
			[]int{1, 3, 5},
			5,
			2,
		},
		{
			[]int{2, 7, 8, 9, 10},
			9,
			3,
		},
	}
	for _, test := range tests {
		output := searchInsert(test.nums, test.target)
		log.Printf("[%d = %d]\n", output, test.expected)
	}
}
