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
			if x+1 < n && y-1 < 0 {
				// 下個點在右邊
				x++
				y = 0
				reverse = true
			} else if x+1 == n {
				// 下個點在下面
				x = n - 1
				y++
				reverse = true
			} else {
				// 正常情況, 右上走
				x++
				y--
			}
		} else {
			if y+1 < m && x-1 < 0 {
				// 下個點在下邊
				y++
				x = 0
				reverse = false
			} else if y+1 == m {
				// 下個點在右邊
				y = m - 1
				x++
				reverse = false
			} else {
				// 正常情況, 左下走
				x--
				y++
			}
		}
		count++
	}

	return result
}

// @lc code=end
