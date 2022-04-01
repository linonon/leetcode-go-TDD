package leetcode

/*
 * @lc app=leetcode.cn id=2024 lang=golang
 *
 * [2024] 考试的最大困扰度
 */

// @lc code=start

type word struct {
	wLen     int
	distance int
	next     *word
}

func maxConsecutiveAnswers(answerKey string, k int) (result int) {

	var t, f = &word{}, &word{}
	var dumt, dumf = &word{}, &word{}
	dumt.next = t
	dumt.next = f
	for i := 0; i < len(answerKey); i++ {
		if answerKey[i] == 'T' {
			t.wLen++
		}
		for ; i < len(answerKey) && answerKey[i] != 'T'; i++ {
			t.distance++
		}
		if i >= len(answerKey) {
			break
		}
		t.next = &word{}
		t = t.next
	}

	return
}

// @lc code=end
