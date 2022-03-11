package pkgu

import "testing"

func TestPrintTreeRootNoe(t *testing.T) {
	r := []int{1, -1, 2, 3, 4, 5, -1, -1, 6, 7, -1, 8, -1, 9, 10, -1, -1, 11, -1, 12, -1, 13, -1, -1, 14}
	ToTreeNode(r)
}
