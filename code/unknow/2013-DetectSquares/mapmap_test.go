package leetcode

import (
	"fmt"
	"testing"
)

func TestMapMapOutput(t *testing.T) {
	var a = make(map[[2]int]int)
	var points = [][2]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {4, 5}}

	for _, v := range points {
		a[v]++
	}

	for xy, Count := range a {
		fmt.Printf("(x:%d,y:%d): Count: %d\n", xy[0], xy[1], Count)
	}

}
