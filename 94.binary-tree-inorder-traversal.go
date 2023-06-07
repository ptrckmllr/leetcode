package main

/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 *
 * https://leetcode.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Easy (74.07%)
 * Likes:    11560
 * Dislikes: 588
 * Total Accepted:    2M
 * Total Submissions: 2.7M
 * Testcase Example:  '[1,null,2,3]'
 *
 * Given the root of a binary tree, return the inorder traversal of its nodes'
 * values.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [1,null,2,3]
 * Output: [1,3,2]
 *
 *
 * Example 2:
 *
 *
 * Input: root = []
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: root = [1]
 * Output: [1]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [0, 100].
 * -100 <= Node.val <= 100
 *
 *
 *
 * Follow up: Recursive solution is trivial, could you do it iteratively?
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}
	current := root
	result = append(result, current.Val)
	for !(current.Left == nil && current.Right == nil && len(stack) == 0) {
		if current.Left != nil {
			// TODO: this will traverse back down the left side once it has gone back up
			stack = append(stack, current)
			current = current.Left
		} else if current.Right != nil {
			stack = append(stack, current)
			current = current.Right
		} else {
			result = append(result, current.Val)
			lastIndex := len(stack) - 1
			current = stack[lastIndex]
			stack = stack[:lastIndex]
		}
	}
	return result
}

// @lc code=end
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	inorderTraversal(root)
}
