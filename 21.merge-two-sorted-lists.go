package main

/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var start ListNode
	var last *ListNode
	last = &start

	for {
		if list1 == nil {
			last.Next = list2
			break
		}
		if list2 == nil {
			last.Next = list1
			break
		}
		if list1.Val > list2.Val {
			last.Next = list2
			last = last.Next
			list2 = list2.Next
		} else {
			last.Next = list1
			last = last.Next
			list1 = list1.Next
		}
	}
	return start.Next
}

// @lc code=end

type ListNode struct {
	Val  int
	Next *ListNode
}
