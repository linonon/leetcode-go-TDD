# 連結兩個 List

## Version 1
龜龜，也太複雜了點。。。
```go
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    // 基礎的前置判斷
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 != nil && list2 == nil {
		return list1
	} else if list1 == nil && list2 != nil {
		return list2
	}

	var x1, x2 *ListNode
    // 記錄 list1 的頭，最後返回的是這個頭
	result := list1
	for x1, x2 = list1, list2; x1 != nil && x2 != nil; {
		if x1.Val <= x2.Val {
            // x1:{ 1->3 } x2:{ 5->7 }
            // 找到 x1 中最後那個小於 x2.Val 的點
			for {
				if x1.Next == nil {
					break
				}
				if x1.Next.Val <= x2.Val {
					x1 = x1.Next
				} else {
					break
				}
			}

            // temp 記錄 x1.Next
			t := x1.Next
            // 效果等於 == 在 x1 和 x1.Next 之間塞一個數
			x1.Next = &ListNode{Val: x2.Val, Next: t}
			if x1.Next.Next == nil {
				x1.Next.Next = x2.Next
				break
			}

            // 更新需要對比點節點位置
            // x1 {1, 1, 2, 4}
            //     ^ --> ^  : x1 = x1.next.next
            // x2 {1, 3, 4}
            // x2  ^->^     : x2 = x2. next
			x1 = x1.Next.Next
			x2 = x2.Next
		} else if x1.Val > x2.Val {
            // 記錄節點位置和值
			t := x1.Next
			v := x1.Val
            // 替換成 x2 的值，然後再後面添加回自己 temp 的值
			x1.Val = x2.Val
			x1.Next = &ListNode{Val: v, Next: t}

            // 更新需要對比的節點位置
            // x1 {5}  
            // x2 {1, 2, 4}

            // x1 {1, 5}    -> x1 {5}
            //     ^->^ : x1 = x1.Next
            // x2 {1, 2, 4} -> x2 {2, 4}
            //     ^->^ : x2 = x2.Next
			x1 = x1.Next
			x2 = x2.Next
		}
	}

	return result
}
```

## Version Most Votes
這也太清晰了
```go
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}
```