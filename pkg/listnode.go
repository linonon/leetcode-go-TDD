package pkgu

import (
	"fmt"

	"github.com/pkg/errors"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	if n == nil {
		return "node == nil"
	}

	var str = ""
	for x, i := n, 0; x != nil && i < 20; x = x.Next {
		if x.Next != nil {
			str += fmt.Sprintf(" %v ->", x.Val)
		} else {
			str += fmt.Sprintf(" %v", x.Val)
		}
		i++
	}

	return str
}

func (n *ListNode) IsEqual(other *ListNode) bool {
	if n == nil && other == nil {
		return true
	} else if n == nil || other == nil {
		return false
	}

	var x1, x2 *ListNode
	for x1, x2 = n, other; x1 != nil && x2 != nil; {
		if x1.Val != x2.Val {
			return false
		}
		x1 = x1.Next
		x2 = x2.Next
	}

	if x1 != nil || x2 != nil {
		return false
	}

	return true
}

func (n *ListNode) Length() (int, error) {
	result := 0
	for ; n != nil; n = n.Next {
		result++
		if result >= 100 {
			return 0, errors.New("this node maybe a ring node")
		}
	}
	return result, nil
}

func ToNode(s []int) *ListNode {
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

func RingNode(s []int) *ListNode {
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

	l1.Next = result
	return result
}
