package pkgu

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func IntSliceToListNode(s []int) *ListNode {
	if len(s) == 0 {
		return nil
	}

	l1 := &ListNode{Val: s[0], Next: nil}
	for i := 1; i < len(s); i++ {
		l2 := &ListNode{Val: s[i], Next: nil}
		l1.Next = l2
		fmt.Println("fmt:", l1.Val)
	}

	return l1
}
