package leetcode

/*
 * @lc app=leetcode.cn id=2013 lang=golang
 *
 * [2013] 检测正方形
 */

// DetectSquares is a struct
// @lc code=start
type DetectSquares struct {
	m map[int]map[int]int
	// {(x,y): 出現次數}
	//   key  <->   value
	// map[x] <-> map[int]int
}

// Constructor likes New()
func Constructor() DetectSquares {
	m := make(map[int]map[int]int)
	return DetectSquares{m}
}

// Add point in x,y into coordinate system.
func (ds *DetectSquares) Add(point []int) {
	x, y := point[0], point[1]
	if ds.m[y] == nil {
		ds.m[y] = map[int]int{}
	}
	ds.m[y][x]++
}

// Count how many right square
//	 ^
//	 |    v(x,col)        v(x+d,col)
//	 |    · ————————————— ·
//	 |				^
//	 |				|
//	 |				d
//	 |				|
//	 |				v
//	 |    · ————————————— ·
//	 |    ^(x,y)          ^(x+d,y)
//	 +———————————————————————————————>
func (ds *DetectSquares) Count(point []int) int {
	x, y := point[0], point[1]
	var result int
	if ds.m[y] == nil {
		return 0
	}

	yCnt := ds.m[y]
	for col, colCnt := range ds.m {
		if col != y {
			d := col - y
			result += colCnt[x] * colCnt[x+d] * yCnt[x+d]
			result += colCnt[x] * colCnt[x-d] * yCnt[x-d]
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
