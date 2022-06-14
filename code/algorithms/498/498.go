package leetcode

/*
 * @lc app=leetcode.cn id=498 lang=golang
 *
 * [498] 对角线遍历
 */

// @lc code=start
func findDiagonalOrder(mat [][]int) []int {
	var (
		m           = len(mat)
		n           = len(mat[0])
		result      = make([]int, 0, m*n)
		x, y    int = 0, 0
		reverse bool
		count   int
	)
	for y < m && x < n && count != n*m {
		result = append(result, mat[y][x])
		if !reverse {
			x++
			y--
			if y < 0 {
				y = 0
				reverse = true
			}
			if x == n {
				y++
				x = n - 1
				reverse = true
			}
		} else {
			x--
			y++
			if x < 0 {
				x = 0
				reverse = false
			}
			if y == m {
				x += 2
				y = m - 1
				reverse = false
			}
		}
		count++
	}

	return result
}

// @lc code=end
