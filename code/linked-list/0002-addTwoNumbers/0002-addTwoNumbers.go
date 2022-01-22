package leetcode

import (
	"fmt"
	. "leetcode/pkg"
)

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var tmp = &ListNode{}
	result := tmp

	// 這裡處理沒有 {9} + {9} = {8,1} 的情況
	for l1 != nil && l2 != nil {

		sum := tmp.Val + l1.Val + l2.Val
		if sum >= 10 {
			tmp.Val = sum - 10
			tmp.Next = &ListNode{Val: 1}
			if l1.Next != nil && l2.Next != nil {
				tmp = tmp.Next
			}
		} else {
			tmp.Val = sum
			if l1.Next != nil || l2.Next != nil {
				tmp.Next = &ListNode{}
				tmp = tmp.Next
			}
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	// 當發現 有超出進位時的情況判斷
	for tmp.Next != nil {
		tmp = tmp.Next
		if l1 != nil {
			sum := l1.Val + tmp.Val
			if sum >= 10 {
				tmp.Val = sum - 10
				tmp.Next = &ListNode{Val: 1}
			} else {
				tmp.Val = sum
				if l1.Next != nil {
					tmp.Next = &ListNode{}
					tmp = tmp.Next
				}
			}
			l1 = l1.Next
		}

		if l2 != nil {
			sum := l2.Val + tmp.Val
			if sum >= 10 {
				tmp.Val = sum - 10
				tmp.Next = &ListNode{Val: 1}
			} else {
				tmp.Val = sum
				if l2.Next != nil {
					tmp.Next = &ListNode{}
					tmp = tmp.Next
				}
			}
			l2 = l2.Next
		}
	}

	// 加完了，沒進位了，把剩下的拼接起來。
	if l1 != nil && tmp.Next == nil {
		tmp.Val = l1.Val
		tmp.Next = l1.Next
	}
	if l2 != nil && tmp.Next == nil {
		tmp.Val = l2.Val
		tmp.Next = l2.Next
	}

	fmt.Println("tmp:", tmp)
	return result
}

// @lc code=end
