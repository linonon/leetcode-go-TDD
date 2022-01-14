# 羅馬數字轉整數

| 字符 | 數值 |
| :--- | ---: |
| I    |    1 |
| V    |    5 |
| X    |   10 |
| L    |   50 |
| C    |  100 |
| D    |  500 |
| M    | 1000 |

通常情況： 大數左， 小數右
特殊情況： 4 & 9 or 40 & 90...: IV = 4, IX = 9, XL = 14, XC = 90...

比較清晰，沒什麼好說的啦
```go
func romanToInt(s string) int {
	romanMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	result := romanMap[string(s[len(s)-1])]

	for i := len(s) - 2; i >= 0; i-- {
		if romanMap[string(s[i])] >= romanMap[string(s[i+1])] {
			result += romanMap[string(s[i])]
		} else {
			result -= romanMap[string(s[i])]
		}
	}

	return result
}
```