package leetcode

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	t := []string{}
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" || string(s[i]) == "[" || string(s[i]) == "{" {
			t = append(t, string(s[i]))
		}
		if string(s[i]) == ")" {
			if !isPopMatch(t, "(") {
				return false
			}
			t = t[:len(t)-1]
		}
		if string(s[i]) == "]" {
			if !isPopMatch(t, "[") {
				return false
			}
			t = t[:len(t)-1]
		}
		if string(s[i]) == "}" {
			if !isPopMatch(t, "{") {
				return false
			}
			t = t[:len(t)-1]
		}
	}

	return len(t) == 0
}

func isPopMatch(t []string, a string) bool {
	if len(t) == 0 {
		return false
	}
	r := t[len(t)-1]
	if r != a || len(t) == 0 {
		return false
	}

	return true
}

// @lc code=end
