# 回文數

## Version1:
```go
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	stringX := strconv.Itoa(x)

	for i := 0; i < len(stringX)/2; i++ {
		if stringX[i] != stringX[len(stringX)-1-i] {
			return false
		}
	}

	return true
}
```

速度太慢了。。。

## Version2
```go
func isPalindrome(x int) bool {
	var reversedNum int
	tmp := x
	for tmp > 0 {
		reversedNum = reversedNum*10 + tmp%10
		tmp = tmp / 10
	}
	return x == reversedNum
}
```
時間/空間複雜度 =O(n), 還可以進步

## Version3
```go
func isPalindrome(x int) bool {
   	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	var reversedNum int
	for x > reversedNum {
		reversedNum = reversedNum*10 + x%10
		x = x / 10
	}

	// 12321:  x = 12, r = 123, r/10 = 12.
	// 1221: x = 12, r = 12
	return x == reversedNum || x == reversedNum/10
}
```

時間複雜度 = O(n/2)
空間複雜度 = O(n)