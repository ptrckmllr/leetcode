package main

/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 *
 * https://leetcode.com/problems/merge-sorted-array/description/
 *
 * algorithms
 * Easy (47.01%)
 * Likes:    10711
 * Dislikes: 1051
 * Total Accepted:    2.1M
 * Total Submissions: 4.4M
 * Testcase Example:  '[1,2,3,0,0,0]\n3\n[2,5,6]\n3'
 *
 * You are given two integer arrays nums1 and nums2, sorted in non-decreasing
 * order, and two integers m and n, representing the number of elements in
 * nums1 and nums2 respectively.
 *
 * Merge nums1 and nums2 into a single array sorted in non-decreasing order.
 *
 * The final sorted array should not be returned by the function, but instead
 * be stored inside the array nums1. To accommodate this, nums1 has a length of
 * m + n, where the first m elements denote the elements that should be merged,
 * and the last n elements are set to 0 and should be ignored. nums2 has a
 * length of n.
 *
 *
 * Example 1:
 *
 *
 * Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
 * Output: [1,2,2,3,5,6]
 * Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
 * The result of the merge is [1,2,2,3,5,6] with the underlined elements coming
 * from nums1.
 *
 *
 * Example 2:
 *
 *
 * Input: nums1 = [1], m = 1, nums2 = [], n = 0
 * Output: [1]
 * Explanation: The arrays we are merging are [1] and [].
 * The result of the merge is [1].
 *
 *
 * Example 3:
 *
 *
 * Input: nums1 = [0], m = 0, nums2 = [1], n = 1
 * Output: [1]
 * Explanation: The arrays we are merging are [] and [1].
 * The result of the merge is [1].
 * Note that because m = 0, there are no elements in nums1. The 0 is only there
 * to ensure the merge result can fit in nums1.
 *
 *
 *
 * Constraints:
 *
 *
 * nums1.length == m + n
 * nums2.length == n
 * 0 <= m, n <= 200
 * 1 <= m + n <= 200
 * -10^9 <= nums1[i], nums2[j] <= 10^9
 *
 *
 *
 * Follow up: Can you come up with an algorithm that runs in O(m + n) time?
 *
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		for i, num := range nums2 {
			nums1[i] = num
		}
	}
	i_m := m - 1
	i_n := n - 1
	for index := m + n - 1; index >= 0; index-- {
		if i_m < 0 {
			nums1[index] = nums2[i_n]
			i_n--
		} else if i_n < 0 {
			nums1[index] = nums1[i_m]
			i_m--
		} else if nums1[i_m] >= nums2[i_n] {
			nums1[index] = nums1[i_m]
			i_m--
		} else {
			nums1[index] = nums2[i_n]
			i_n--
		}
	}
}

// @lc code=end

func main() {
	merge([]int{2, 0}, 1, []int{1}, 1)
}
