package leetcode

import (
	"fmt"
	. "leetcode/pkg"
	"testing"
)

func Test589(t *testing.T) {
	var tmp = []int{1, -1, 2, 3, 4, 5, -1, -1, 6, 7, -1, 8, -1, 9, 10, -1, -1, 11, -1, 12, -1, 13, -1, -1, 14}
	var tmpw = []int{1, -1, 3, 2, 4, -1, 5, 6}
	n := ToTreeNode(tmp)
	fmt.Println(preorder(n))
	n2 := ToTreeNode(tmpw)
	fmt.Println(preorder(n2))
}
