package leetcode

/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N 叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
import t "leetcode/pkg"

func preorder(root *t.NTreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result = make([]int, 0)

	f(root, &result)

	return result
}

func f(root *t.NTreeNode, result *[]int) {
	*result = append(*result, root.Val)
	if len(root.Children) == 0 {
		return
	} else {
		for _, child := range root.Children {
			f(child, result)
		}
	}
}

// @lc code=end
