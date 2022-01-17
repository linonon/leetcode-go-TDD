
# 有效的括号

## Version 1
```go
func isValid(s string) bool {
	t := []string{}
	r := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" || string(s[i]) == "[" || string(s[i]) == "{" {
			t = append(t, string(s[i]))
		}
		if string(s[i]) == ")" {
			if len(t) == 0 {
				return false
			}
			r = t[len(t)-1]
			if r != "(" || len(t) == 0 {
				return false
			}
			t = t[:len(t)-1]
		}
		if string(s[i]) == "]" {
			if len(t) == 0 {
				return false
			}
			r = t[len(t)-1]
			if r != "[" || len(t) == 0 {
				return false
			}
			t = t[:len(t)-1]
		}
		if string(s[i]) == "}" {
			if len(t) == 0 {
				return false
			}
			r = t[len(t)-1]
			if r != "{" || len(t) == 0 {
				return false
			}
			t = t[:len(t)-1]
		}
	}

	return len(t) <= 0
}
```
- 思路： 碰到左邊的，就push進數組，碰到右邊的，就從數組中pop一個出來，看看右邊的和左邊的是不是一種類型的。不是就返回false
- 總結：這個方法簡直不要太麻煩咯。

## Version 1 (Refeactoring)
```go
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
```

## Version Most Voted

思路一樣的
```go
func isValid(s string) bool {
	tr := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	stack := make([]rune, 0, 1)

	for _, ch := range s {
		switch ch {
		case 40, 123, 91: // "([{"
			stack = append(stack, ch)
		case 41, 125, 93: // ")]}"
			if len(stack) == 0 || stack[len(stack)-1] != tr[ch] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}
```