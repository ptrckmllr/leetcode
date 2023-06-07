package main

/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 || x == 2 || x == 3 {
		return 1
	}

	low, high := 0, x/2+1

	for low < high {
		mid := (high + low) / 2
		if low == mid {
			return mid
		}
		if x < mid*mid {
			high = mid
		} else if x > mid*mid {
			low = mid
		} else {
			return mid
		}
	}
	return low
}

// @lc code=end

func main() {
	println(mySqrt(8))
}
