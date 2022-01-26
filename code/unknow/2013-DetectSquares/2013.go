package leetcode

/*
 * @lc app=leetcode.cn id=2013 lang=golang
 *
 * [2013] 检测正方形
 */

// DetectSquares is a struct
// @lc code=start
type DetectSquares struct {
	points [][]int
}

// Constructor likes New()
func Constructor() DetectSquares {
	points := make([][]int, 0)
	return DetectSquares{points}
}

// Add point in x,y into coordinate system.
func (ds *DetectSquares) Add(point []int) {
	ds.points = append(ds.points, point)
}

// Count how many ways to get a square
func (ds *DetectSquares) Count(point []int) int {
	var absolute = func(a, b int) int {
		if a >= b {
			return a - b
		}
		return b - a
	}
	const x, y int = 0, 1
	result := 0
	for _, p1 := range ds.points {
		// x Equal
		if p1[x] == point[x] {
			for _, p2 := range ds.points {
				// y Equal
				if p2[y] == point[y] {
					tx := p2[x]
					ty := p1[y]
					if absolute(point[y], ty) != absolute(point[x], tx) {
						continue
					}

					for k := 0; k < len(ds.points); k++ {
						if ds.points[k][x] == tx && ds.points[k][y] == ty {
							result++
						}
					}
				}
			}
		}
	}

	return result
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */
// @lc code=end
