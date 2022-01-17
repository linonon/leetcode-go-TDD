package pkgu

type ListNode struct {
	Val  int
	Next *ListNode
}

func IntSliceToListNode(s []int) *ListNode {
	if len(s) == 0 {
		return nil
	}

	l1 := &ListNode{Val: s[0], Next: nil}
	result := l1
	for i := 1; i < len(s); i++ {
		l2 := &ListNode{Val: s[i], Next: nil}
		l1.Next = l2
		l1 = l1.Next
	}

	return result
}
