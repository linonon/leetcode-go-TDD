package leetcode

import (
	"fmt"
	. "leetcode/pkg"
)

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var x1, x2 *ListNode
	for x1 = list1; x1 != nil; x1 = x1.Next {
		fmt.Println(x1.Val)
	}

	for x2 = list2; x2 != nil; x2 = x2.Next {
		fmt.Println(x2.Val)
	}

	return list1
}

// @lc code=end
