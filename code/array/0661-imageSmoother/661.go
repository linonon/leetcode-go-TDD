package leetcode

/*
 * @lc app=leetcode.cn id=661 lang=golang
 *
 * [661] 图片平滑器
 */

// @lc code=start
func imageSmoother(img [][]int) (result [][]int) {
	m, n := len(img), len(img[0])
	result = make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
	}
	if n == 1 && m == 1 {
		return img
	}
	if n == 1 && m >= 2 {
		for y := 0; y < m; y++ {
			if y == 0 {
				result[y][0] = (img[y][0] + img[y+1][0]) / 2
			} else if y == m-1 {
				result[y][0] = (img[y-1][0] + img[y][0]) / 2
			} else {
				result[y][0] = (img[y-1][0] + img[y][0] + img[y+1][0]) / 3
			}
		}
		return
	}

	if m == 1 && n >= 2 {
		for x := 0; x < n; x++ {
			if x == 0 {
				result[0][x] = (img[0][x] + img[0][x+1]) / 2
			} else if x == n-1 {
				result[0][x] = (img[0][x-1] + img[0][x]) / 2
			} else {
				result[0][x] = (img[0][x-1] + img[0][x] + img[0][x+1]) / 3
			}
		}
		return
	}

	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			result[y][x] = check(x, y, n, m, img)
		}
	}
	return
}

func check(x, y, n, m int, img [][]int) int {
	if 1 <= x && x <= n-2 && 1 <= y && y <= m-2 {
		return (img[y-1][x-1] + img[y-1][x] + img[y-1][x+1] + img[y][x-1] + img[y][x] + img[y][x+1] + img[y+1][x-1] + img[y+1][x] + img[y+1][x+1]) / 9
	}
	if x == 0 && y == 0 {
		return (img[y][x] + img[y][x+1] + img[y+1][x] + img[y+1][x+1]) / 4
	}
	if x == 0 && y == m-1 {
		return (img[y-1][x] + img[y-1][x+1] + img[y][x] + img[y][x+1]) / 4
	}
	if x == n-1 && y == 0 {
		return (img[y][x-1] + img[y][x] + img[y+1][x-1] + img[y+1][x]) / 4
	}
	if x == n-1 && y == m-1 {
		return (img[y-1][x-1] + img[y-1][x] + img[y][x-1] + img[y][x]) / 4
	}
	if x == 0 && 1 <= y && y <= m-2 {
		return (img[y-1][x] + img[y-1][x+1] + img[y][x] + img[y][x+1] + img[y+1][x] + img[y+1][x+1]) / 6
	}
	if y == 0 && 1 <= x && x <= n-2 {
		return (img[y][x-1] + img[y][x] + img[y][x+1] + img[y+1][x-1] + img[y+1][x] + img[y+1][x+1]) / 6
	}
	if x == n-1 && 1 <= y && y <= m-2 {
		return (img[y-1][x-1] + img[y-1][x] + img[y][x-1] + img[y][x] + img[y+1][x-1] + img[y+1][x]) / 6
	}
	if y == m-1 && 1 <= x && x <= n-2 {
		return (img[y-1][x-1] + img[y-1][x] + img[y-1][x+1] + img[y][x-1] + img[y][x] + img[y][x+1]) / 6
	}

	return 0
}

// @lc code=end
