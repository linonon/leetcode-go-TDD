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
	if list1 == nil || list2 == nil {
		return nil
	}

	var x1, x2 *ListNode
	x1 = list1
	for i := 0; i < 20; i++ {
		fmt.Println("x1.Val:", x1.Val)
		if x1.Next != nil {
			x1 = x1.Next
		} else {
			break
		}
	}

	for x2 = list2; x2 != nil; x2 = x2.Next {
		fmt.Println(x2.Val)
	}

	return list1
}

// @lc code=end
